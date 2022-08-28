/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a ComplaintRequest, ComplaintResult
type SmartContract struct {
	contractapi.Contract
}

/* 민원요청*/
type ComplaintRequest struct {
	RequestId         string `json:"request_id"`         //신청번호
	UserId            string `json:"user_id"`              //UserId
	RequesterName     string `json:"requester_name"`     //신청인이름
	PhoneNumber       string `json:"phone_number"`       //연락처
	Address           string `json:"address"`            //주소
	Open              string `json:"open"`               //민원 공개 여부
	Title             string `json:"title"`              //민원 제목
	Content           string `json:"content"`            //민원 내용
	ComplaintLocation string `json:"complaint_location"` //민원 발생지역
	RequestDate       string `json:"request_date"`       //신청일
	ReceptionStatus   int    `json:"reception_status"`   //접수상태
	ReceptionDate     string `json:"reception_date"`     //접수일
	Result	  *ComplaintResult `json:complaint_result`	//민원결과
}


/* 민원결과*/
type ComplaintResult struct {
	RequestId	 string `json:request_id`		//신청번호
	ResultId     string `json:"result_id"`     //결과번호
	UserId        string `json:"user_id"`        //UserId
	Agency        string `json:"agency"`         //처리기관
	Manager       string `json:"manager"`        //담당자
	ResultDate    string `json:"result_date"`    //답변일
	ResultContent string `json:"result_content"` //처리결과(답변내용)
}

type User struct {
	UserId string `json:"userId"`	//사용자명
	UserType int `json:"userType"`
	Password string `json:"password"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *ComplaintRequest
}

// 사용자 추가
func (s *SmartContract) AddUser (ctx contractapi.TransactionContextInterface, userId string, password string) (string, error) {
	user := User{
		UserId: userId,
		Password: password,
	}

	userAsBytes, _ := json.Marshal(user)

	return userId, ctx.GetStub().PutState(userId, userAsBytes)
}

// CreateComplaintRequest adds a new CreateComplaintRequest to the world state with given details
func (s *SmartContract) AddComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string, userId string, requesterName string, phoneNumber string, address string, open string, title string, content string, complaintLocation string, requestDate string) (string, error) {
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
func (s *SmartContract) GetComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string) (*ComplaintRequest, error) {
	complaintRequestAsBytes, err := ctx.GetStub().GetState(requestId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if complaintRequestAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", requestId)
	}

	complaintRequest := new(ComplaintRequest)
	_ = json.Unmarshal(complaintRequestAsBytes, complaintRequest)

	return complaintRequest, nil
}

//민원 수정 who.백성
func (s *SmartContract) UpdateComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string, userId string, phoneNumber string, address string, open string, title string, content string, complaintLocation string) (string, error) {

	complaintRequest, err := s.GetComplaintRequest(ctx, requestId);

	if err != nil {
		return requestId, err
	}

	complaintRequest.PhoneNumber = phoneNumber
	complaintRequest.Address = address
	complaintRequest.Open = open
	complaintRequest.Title = title
	complaintRequest.Content = content
	complaintRequest.ComplaintLocation = complaintLocation

	complaintRequestAsBytes, _ := json.Marshal(complaintRequest)
	
	return requestId, ctx.GetStub().PutState(requestId, complaintRequestAsBytes)
}

//민원 접수 상태 수정 who.관아
func (s *SmartContract) UpdateComplaintRequestStatus (ctx contractapi.TransactionContextInterface, requestId string, receptionStatus int, receptionDate string) (string, error) {
	complaintRequest, err := s.GetComplaintRequest(ctx, requestId);

	if err != nil {
		return requestId, err
	}

	complaintRequest.ReceptionStatus = receptionStatus
	complaintRequest.ReceptionDate = receptionDate

	complaintRequestAsBytes, _ := json.Marshal(complaintRequest)
	
	return requestId, ctx.GetStub().PutState(requestId, complaintRequestAsBytes)
}

//민원 요청 삭제 who.백성
func (s *SmartContract) DeleteComplaintRequest (ctx contractapi.TransactionContextInterface, requestId string) (string, error) {

	complaintRequest, err := s.GetComplaintRequest(ctx, requestId)
	if err != nil {
		return requestId, err
	}
	if complaintRequest == nil {
		return requestId, fmt.Errorf("the ComplaintRequest %s does not exist", requestId)
	}

	return requestId, ctx.GetStub().DelState(requestId)

}

//나의 민원 리스트 who.백성
func (s *SmartContract) ListComplaintRequestByUserId (ctx contractapi.TransactionContextInterface, userId string) ([]ComplaintRequest, error) {
	return []ComplaintRequest{}, nil
}

//유사민원검색 who.백성
func (s *SmartContract) SearchComplaintRequestByTitle (ctx contractapi.TransactionContextInterface, title string) ([]ComplaintRequest, error) {
	return []ComplaintRequest{}, nil
}

//전체 민원 리스트 who.관원
func (s *SmartContract) ListComplaintRequestAll (ctx contractapi.TransactionContextInterface) ([]*ComplaintRequest, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var complaintRequests []*ComplaintRequest
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var complaintRequest ComplaintRequest
		err = json.Unmarshal(queryResponse.Value, &complaintRequest)
		if err != nil {
			return nil, err
		}
		complaintRequests = append(complaintRequests, &complaintRequest)
	}

	return complaintRequests, nil

}

//민원 결과 발행 who.관원
func (s *SmartContract) AddComplaintResult (ctx contractapi.TransactionContextInterface, requestId string, resultId string, agency string, userId string, manager string, resultDate string, resultContent string) (string, error) {
//	return resultId, nil

	complaintResult := ComplaintResult{
		RequestId:	requestId,
		ResultId:         	resultId,
		Agency:     		agency,
		UserId:           userId,
		Manager:       manager,
		ResultDate:           resultDate,
		ResultContent:              resultContent,
	}

	complaintRequest, err := s.GetComplaintRequest(ctx, requestId);

	if err != nil {
		return requestId, err
	}

	complaintRequest.Result = &complaintResult

	complaintRequestAsBytes, _ := json.Marshal(complaintRequest)
	
	return resultId, ctx.GetStub().PutState(requestId, complaintRequestAsBytes)
}

//민원 결과 수정 who.관원
func (s *SmartContract) UpdateComplaintResult (ctx contractapi.TransactionContextInterface, requestId string, agency string, userId string, manager string, resultDate string, resultContent string) (string, error) {
	complaintRequest, err := s.GetComplaintRequest(ctx, requestId);

	if err != nil {
		return requestId, err
	}

	complaintResult := complaintRequest.Result

	complaintResult.Agency = agency;
	complaintResult.UserId = userId;
	complaintResult.Manager = manager;
	complaintResult.ResultDate = resultDate;
	complaintResult.ResultContent = resultContent;
	

	complaintRequestAsBytes, _ := json.Marshal(complaintRequest)
	
	return requestId, ctx.GetStub().PutState(requestId, complaintRequestAsBytes)

}

//민원 결과 번호
func (s *SmartContract) GetComplaintResult (ctx contractapi.TransactionContextInterface, resultId string) (*ComplaintResult, error) {
	return &ComplaintResult{}, nil
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
