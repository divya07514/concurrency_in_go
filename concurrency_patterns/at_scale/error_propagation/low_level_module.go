package errorpropagation

import "os"

type LowLevelErr struct {
	error
}

func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(WrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
