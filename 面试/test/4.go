var (
    MaxWorker = os.Getenv("MAX_WORKERS")
    MaxQueue  = os.Getenv("MAX_QUEUE")
)


dispatcher := NewDispatcher(MaxWorker)
dispatcher.Run()