package main

func main() {

    infoCh := make(chan []VideoInfoPair)
    //downloadCh := make(chan []DownloadedPair) 

    go GetVideoPairInfo(infoCh)
    //go VideoDownloaderRoutine(infoCh, downloadCh)
    //go VideoProcessingRoutine(downloadCh)

}
