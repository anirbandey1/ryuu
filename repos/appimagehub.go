package repos

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


type AppimageHubOCS struct {
    XMLName xml.Name `xml:"ocs"`
    Meta struct {
        Status string `xml:"status"`
        Statuscode int8 `xml:"statuscode"`
        Message string `xml:"message"`
        TotalItems int8 `xml:"totalitems"`
        ItemsPerPage int8 `xml:"itemsperpage"`
    } `xml:"meta"`
    Data struct {
        Content []Content `xml:"content"`

    } `xml:"data"`
}


type Content struct {
    Details string `xml:"details,attr"`
    Id string `xml:"id"`
    Name string `xml:"name"`
    Version string `xml:"version"`
    TypeId string `xml:"typeid"`
    Created string `xml:"created"`
    Changed string `xml:"changed"`
    Description string `xml:"description"`
    Tags string `xml:"tags"`
    DownloadWay1 string `xml:"downloadway1"`
    DownloadType1 string `xml:"downloadtype1"`
    DownloadPrice1 string `xml:"downloadprice1"`
    DownloadLink1 string `xml:"downloadlink1"`
    DownloadName1 string `xml:"downloadname1"`
    DownloadSize1 string `xml:"downloadsize1"`
    DownloadGpgFingerprint1 string `xml:"downloadgpgfingerprint1"`
    DownloadGpgSignature1 string `xml:"downloadgpgsignature1"`
    DownloadPackageName1 string `xml:"downloadpackagename1"`
    DownloadRepository1 string `xml:"downloadrepository1"`
    DownloadPackageType1 string `xml:"download_package_type1"`
    DownloadPackageArch1 string `xml:"download_package_arch1"`
    DownloadTags1 string `xml:"downloadtags1"`
    DownloadMDSum1 string `xml:"downloadmd5sum1"`
    DownloadWay2 string `xml:"downloadway2"`
    DownloadType2 string `xml:"downloadtype2"`
    DownloadPrice2 string `xml:"downloadprice2"`
    DownloadLink2 string `xml:"downloadlink2"`
    DownloadName2 string `xml:"downloadname2"`
    DownloadSize2 string `xml:"downloadsize2"`
    DownloadGpgFingerprint2 string `xml:"downloadgpgfingerprint2"`
    DownloadGpgSignature2 string `xml:"downloadgpgsignature2"`
    DownloadPackageName2 string `xml:"downloadpackagename2"`
    DownloadRepository2 string `xml:"downloadrepository2"`
    DownloadPackageType2 string `xml:"download_package_type2"`
    DownloadPackageArch2 string `xml:"download_package_arch2"`
    DownloadTags2 string `xml:"downloadtags2"`
    DownloadMDSum2 string `xml:"downloadmd5sum2"`


}

func SearchAppimage(target string){

    requestUrl  := fmt.Sprintf("https://www.appimagehub.com/ocs/v1/content/data?search=%s",target)

    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    request, err := http.NewRequest("GET",requestUrl,nil)

    if err != nil {
        log.Fatal()

    }


    request.Header.Set("User-Agent","Not firefox")


    // response, err := http.Get(fmt.Sprintf("https://www.appimagehub.com/ocs/v1/content/data?search=%s",target)) 
    response, err := client.Do(request)

    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal(err)

    }


    // fmt.Println(string(body))
    ocs := AppimageHubOCS{}
    xml.Unmarshal(body, &ocs)

    fmt.Println(ocs)





}
