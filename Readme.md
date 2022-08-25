# Reporting in Fabric

## 개요 
블록체인 민원 프로젝트
- Network 구조 : https://docs.google.com/presentation/d/1Uw4QljW7FFWONqQ7pinJK1JejxUvLGf-xW37tvZsmL4/edit#slide=id.g14a3fd3ec92_1_0

https://go.dev/play/p/LxJA3WMXdgy

<hr/>

## 폴더구조

- application
- contract
- network : 네트워크 구성 (Network 구조 그림 참고) 
  - startnetwork.sh     
  - createchannel.sh
  - setAnchorPeerUpdate.sh
  - networkdown.sh

<hr/>

## GIT 

https://github.com/gymnopedy01/fraudreporting

```sh
$> git init 
$> git add --all
$> git commit 
$> git config --global user.name "th"
$> git branch -M main
$> git remote add origin https://github.com/gymnopedy01/fraudreporting.git
$> git push -u origin main
```

```sh
$> git config --global core.editor vi
```

```sh
$> git add --all
$> git commit
$> git push
```


## Chain Code

shim API

https://pkg.go.dev/github.com/hyperledger/fabric-chaincode-go@v0.0.0-20220720122508-9207360bbddd/shim#ChaincodeStub.GetQueryResult



```go
type ComplaintRequest struct {
	RequestID         string `json:"request_id"`         //신청번호
	RequesterName     string `json:"requester_name"`     //신청인이름
	PhoneNumber       string `json:"phone_number"`       //연락처
	Address           string `json:"address"`            //주소
	Open              bool   `json:"open"`               //민원 공개
	Title             string `json:"title"`              //민원 제목
	Content           string `json:"content"`            //민원 내용
	ComplaintLocation string `json:"complaint_location"` //민원 발생지역
	ReceptionStatus   int    `json:"reception_status"`   //접수상태
	RequestDate       string `json:"request_date"`       //신청일


type ComplaintRequest struct {
RequestID         string //신청번호
RequesterName     string //신청인이름
PhoneNumber       string //연락처
Address           string //주소
Open              bool   //민원 공개
Title             string //민원 제목
Content           string //민원 내용
ComplaintLocation string //민원 발생지역
ReceptionStatus   int    //접수상태
RequestDate       string //신청일
}

type ComplaintResult struct {
Agency        int    //처리기관
RequestID     string //접수번호
ReceptionDate string //접수일
Manager       string //담당자
ResultDate    string //답변일
ResultContent string //처리결과(답변내용)
}
```