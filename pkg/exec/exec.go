package exec

import (
	"errors"
	"fmt"
	"github.com/katainaka0503/berglas-aws/pkg/resolution"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

var (
	stdin  = os.Stdin
	stdout = os.Stdout
	stderr = os.Stderr

	misuseExitCode = 61
	fetchExitCode  = 60
)

func Exec(command string, args []string) error {
	resolver, err := resolution.NewResolverWithContext()
	if err != nil {
		return misuseError(err)
	}

	env := os.Environ()

	for i, value := range env {
		keyAndValue := strings.SplitN(value, "=", 2)
		if len(keyAndValue) < 2 {
			continue
		}

		name, value := keyAndValue[0], keyAndValue[1]

		if resolution.IsResolvable(value) {
			continue
		}

		fetchedValue, err := resolver.Resolve(value)
		if err != nil {
			return fetchError(fmt.Errorf("failed to fetch value of %v=%v: %w", name, value, err))
		}

		env[i] = fmt.Sprintf("%v=%v", name, fetchedValue)
	}

	cmd := exec.Command(command, args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = env

	if err := cmd.Start(); err != nil {
		err = fmt.Errorf("failed to exec command \"%v\": %w", strings.Join(append([]string{command}, args...), " "), err)
		return misuseError(err)
	}

	doneCh := make(chan struct{}, 1)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)

	go func() {
		for {
			select {
			case sig := <-signalCh:
				if cmd.Process == nil {
					return
				}
				if signalErr := cmd.Process.Signal(sig); signalErr != nil && err == nil {
					fmt.Fprintf(stderr, "Failed to signal command: %v\n", err)
				}
			case <-doneCh:
				signal.Reset()
				close(signalCh)
			}

		}
	}()

	if err := cmd.Wait(); err != nil {
		close(doneCh)
		var exitErr *exec.ExitError
		if ok := errors.As(err, &exitErr); ok && exitErr.ProcessState != nil {
			code := exitErr.ProcessState.ExitCode()
			return exitWithCode(code, fmt.Errorf("command exited with non-zero code %v: %w", code, exitErr))
		}
	}

	return nil
}

type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string {
	return fmt.Sprintf("exit with code %v: %v", e.Code, e.Err)
}

func (e *ExitError) UnWrap() error {
	return e.Err
}

func exitWithCode(code int, err error) error {
	return &ExitError{
		Code: code,
		Err:  err,
	}
}

func misuseError(err error) error {
	return exitWithCode(misuseExitCode, err)
}

func fetchError(err error) error {
	return exitWithCode(fetchExitCode, err)
}
