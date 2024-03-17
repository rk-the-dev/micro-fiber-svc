package cmd

func Run() {
	svc := NewSVC("payroll-svc", 3003)
	svc.Start()
}
