package rpc_lib

type Watcher int

func (w *Watcher) GetInfo( arg int, result *int ) error{
    *result = 1234
    return nil
}


