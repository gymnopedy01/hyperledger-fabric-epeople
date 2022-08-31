# Civil Complaint handling system On Hyperledger Fabric Blockchain

## 개요 
 민원 처리 시스템 on 하이퍼렛져 패브릭
- 프로젝트명 : EPEOPLE
- [프레젠테이션](https://docs.google.com/presentation/d/1Uw4QljW7FFWONqQ7pinJK1JejxUvLGf-xW37tvZsmL4/edit)
- [스마트컨트랙트.sample](https://go.dev/play/p/LoiCn3Cvqv4)

<hr/>

## 폴더구조
- application : 
  - server.js 서버 (node.js Express)
  - views : html 파일들
  - startEPP.sh (fabric 한방 구동)
- contract : 체인코드 (GO)
- network : fabric network 설정 
  - startnetwork.sh     
  - createchannel.sh
  - setAnchorPeerUpdate.sh
  - networkdown.sh

<hr/>

## GIT 

https://github.com/gymnopedy01/epeople

```sh
$> git init 
$> git add --all
$> git commit 
$> git config --global user.name "th"
$> git branch -M main
$> git remote add origin https://github.com/gymnopedy01/epeople.git
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


## Contract (Chain Code)

```sh
$> go mod init epeople
$> go mod tidy
$> go build epeople.go
```

### shim API
- https://pkg.go.dev/github.com/hyperledger/fabric-chaincode-go@v0.0.0-20220720122508-9207360bbddd/shim#ChaincodeStub.GetQueryResult

### rich query examples
- https://github.com/hyperledger/fabric-samples/blob/main/asset-transfer-ledger-queries/chaincode-go/asset_transfer_ledger_chaincode.go

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

## Application 설치
```sh
$> git clone https://github.com/gymnopedy01/epeople.git
$> cd application
$> npm install
$> node enrollAdmin.js			--어드민 생성
$> node registerUser.js			--appUser 생성
```

## Applicaton 서버 구동

네트워크 구동후 서버를 구동시킵니다.
```sh
$> cd application
$> ./startEPP.sh
$> node server.js
```

