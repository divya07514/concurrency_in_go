package errorpropagation

import "os/exec"

type IntermediateErr struct {
	error
}

func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := IsGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{WrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)}
	} else if !isExecutable {
		return WrapError(
			nil,
			id,
			"cannot run job %q: requisite binaries are not executable",
		)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}
