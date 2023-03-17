package main

import (
	"os/exec"
)

type step struct {
	name string
	exe string
	args []string 
	message string 
	proj string 
}

func newStep(name, exe, message, proj string, args []string) step {
	return step{
		name: name,
		exe: exe,
		message: message,
		args: args,
		proj: proj,
	}
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step: s.name,
			msg: "failed to execute",
			cause: err,
		}
	}
	
	// if s.name == "git commit" {
	// 	output := strings.Split(out.String(), "")
	// 	commitSha := strings.Join(output[8:15], "")
	// 	env := strings.Split(os.Getenv("APP_VERSION"), "")
	// 	versionNum, _ := strconv.Atoi(strings.Join(env[4:6], "")) 

	// 	versionNum++
	// 	newVersionNum := strconv.Itoa(versionNum)
	// 	file, err := os.OpenFile(".env", os.O_RDWR, 0644)

	// 	if err != nil {
	// 		log.Fatalf("failed opening file: %s", err)
	// 	}
	// 	defer file.Close()

	// 	// edit this depending on your .env
	// 	_, err = file.WriteAt([]byte(newVersionNum), 18)

	// 	if err != nil {
	// 		log.Fatalf("failed writing to file: %s", err)
	// 	}

	// 	_, err = file.WriteAt([]byte(commitSha), 28)

	// 	if err != nil {
	// 		log.Fatalf("failed writing to file: %s", err)
	// 	}
	// }

	// if s.name == "Generating static files" {
	// 	for {
	// 		output, _, _ := buf.ReadLine()
	// 		fmt.Fprintln(os.Stdout, string(output))
	// 	}
		
	// }

	return s.message, nil
}