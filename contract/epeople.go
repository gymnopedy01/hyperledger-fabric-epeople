/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}



/* 민원요청*/
type ComplaintRequest struct {
	RequestId         string `json:"request_id"`         //신청번호
	UserId            string `json:user_id`              //UserId
	RequesterName     string `json:"requester_name"`     //신청인이름
	PhoneNumber       string `json:"phone_number"`       //연락처
	Address           string `json:"address"`            //주소
	Open              bool   `json:"open"`               //민원 공개 여부
	Title             string `json:"title"`              //민원 제목
	Content           string `json:"content"`            //민원 내용
	ComplaintLocation string `json:"complaint_location"` //민원 발생지역
	RequestDate       string `json:"request_date"`       //신청일
	ReceptionStatus   int    `json:"reception_status"`   //접수상태
	ReceptionDate     string `json:"reception_date"`     //접수일
	Result	  ComplaintResult `json:complate_result`	//민원결과
}


/* 민원결과*/
type ComplaintResult struct {
	ResultId     string `json:"people_name"`     //결과번호
	UserId        string `json:"user_id"`        //UserId
	Agency        int    `json:"agency"`         //처리기관
	Manager       string `json:"manager"`        //담당자
	ResultDate    string `json:"result_date"`    //답변일
	ResultContent string `json:"result_content"` //처리결과(답변내용)
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *ComplaintRequest
}

// 사용자 추가
func (s *SmartContract) AddUser (ctx contractapi.TransactionContextInterface, userId string, userType int) (string, error) {
	return "", nil
}

// CreateComplaintRequest adds a new CreateComplaintRequest to the world state with given details
func (s *SmartContract) CreateComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string, userId string, requesterName string, phoneNumber string, address string, open bool, title string, content string, complaintLocation string, requestDate string) (string, error) {
	complaintRequest := ComplaintRequest{
		RequestId:         requestId,
		UserId:            userId,
		RequesterName:     requesterName,
		PhoneNumber:       phoneNumber,
		Address:           address,
		Open:              open,
		Title:             title,
		Content:           content,
		ComplaintLocation: complaintLocation,
		RequestDate:       requestDate,
		ReceptionStatus:   0,
	}

	complaintRequestAsBytes, _ := json.Marshal(complaintRequest)

	return userId, ctx.GetStub().PutState(requestId, complaintRequestAsBytes)
}

//민원 조회 who.백성
func (s *SmartContract) GetComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string) (ComplaintRequest, error) {
	return ComplaintRequest{}, nil
}

//민원 수정 who.백성
func (s *SmartContract) UpdateComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string, userId string, phoneNumber string, address string, open bool, title string, content string, complaintLocation string) (string, error) {
	return requestId, nil
}

//민원 접수 상태 수정 who.관아
func (s *SmartContract) UpdateComplaintRequestStatus (ctx contractapi.TransactionContextInterface, requestId string, receptionStatus int, receptionDate string) (string, error) {
	return requestId, nil
}

//민원 요청 삭제 who.백성
func (s *SmartContract) DeleteComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string) (string, error) {
	return requestId, nil
}

//나의 민원 리스트
func (s *SmartContract) ListComplaintRequestByUserId (ctx contractapi.TransactionContextInterface, userId string) ([]ComplaintRequest, error) {
	return []ComplaintRequest{}, nil
}

//유사민원검색 who.백성
func (s *SmartContract) searchComplaintRequestByTitle (ctx contractapi.TransactionContextInterface, title string) ([]ComplaintRequest, error) {
	return []ComplaintRequest{}, nil
}

//전체 민원 리스트 who.관원
func (s *SmartContract) searchComplaintRequestByTitle (ctx contractapi.TransactionContextInterface) ([]ComplaintRequest, error) {
	return []ComplaintRequest{}, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
