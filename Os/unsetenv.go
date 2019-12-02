package main

import "os"

func  main()  {
	os.Setenv("TEMPDIR","/my/temp")
	defer  os.Unsetenv("TMPDIR")
}
