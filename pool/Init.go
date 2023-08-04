package pool

func init() {
	//go func() {
	//	quit := make(chan os.Signal)
	//	signal.Notify(quit, syscall.SIGINT, syscall.SIGINT)
	//	for {
	//		select {
	//		case <-quit:
	//		default:
	//			log.Print(" go-kit init")
	//			time.Sleep(100)
	//
	//		}
	//
	//	}
	//
	//}()

}

func DefaultStart() *GoesPool {
	goes := NewSetGoesPool(10)
	goes.Start()
	defer goes.Shutdown()

	return goes
}

func defaultPool() {

}
