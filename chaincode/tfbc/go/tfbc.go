/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define Structure for Insurance Claim
type InsuranceClaim struct {
	ClaimID			string		`json:"ClaimID"`
	PolicyNumber    string		`json:"PolicyNumber"`
	EntryDate		string		`json:"EntryDate"`
	InsuranceCompany    	string   	`json:"InsuranceCompany"`
	PlaceOfService		string		`json:"PlaceOfService"`
	ProviderName		string		`json:"ProviderName"`
	ClaimAmount		int		`json:"ClaimAmount"`
	DateOfService 	string 	`json:"DateOfService"`
	DiagnosCode 	string 	`json:"DiagnosCode"`
	ProcedureCode 	string 	`json:"ProcedureCode"`
	TypeOfService  string 	`json:"TypeOfService"`
	ClaimStatus		string		`json:"ClaimStatus"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "requestClaim" {
		return s.requestClaim(APIstub, args)
	} else if function == "acceptClaim" {
		return s.acceptClaim(APIstub, args)
	} else if function == "adjudicateClaim" {
		return s.adjudicateClaim(APIstub, args)
	} else if function == "approveClaim" {
		return s.approveClaim(APIstub, args)
	} else if function == "rejectClaim" {
		return s.rejectClaim(APIstub, args)
	}else if function == "getClaim" {
		return s.getClaim(APIstub, args)
	}else if function == "getClaimHistory" {
		return s.getClaimHistory(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}





// This function is initiated by Patient 
func (s *SmartContract) requestClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

/*
(0)	ClaimID			string		`json:"ClaimID"`
(1)	PolicyNumber    string		`json:"PolicyNumber"`
(2)	EntryDate		string		`json:"EntryDate"`
(3)	InsuranceCompany    	string   	`json:"InsuranceCompany"`
(4)	PlaceOfService		string		`json:"PlaceOfService"`
(5)	ProviderName		string		`json:"ProviderName"`
(6)	ClaimAmount		int		`json:"ClaimAmount"`
(7)	DateOfService 	string 	`json:"DateOfService"`
(8)	DiagnosCode 	string 	`json:"DiagnosCode"`
(9)	ProcedureCode 	string 	`json:"ProcedureCode"`
(10)	TypeOfService  string 	`json:"TypeOfService"`
(11)	ClaimStatus		string		`json:"ClaimStatus"`
*/
	
	ClaimID := args[0];
	PolicyNumber := args[1];
	EntryDate := args[2];
	InsuranceCompany := args[3];
	PlaceOfService := args[4];
	ProviderName := args[5];
	ClaimAmount, err := strconv.Atoi(args[6]);
	DateOfService := args[7];
	DiagnosCode := args[8];
	ProcedureCode := args[9];
	TypeOfService := args[10];
	if err != nil {
		return shim.Error("Not able to parse Claim Amount")
	}


	LC := InsuranceClaim{ClaimID: ClaimID, PolicyNumber: PolicyNumber, EntryDate: EntryDate, InsuranceCompany: InsuranceCompany, PlaceOfService: PlaceOfService,	ProviderName: ProviderName, ClaimAmount: ClaimAmount, DateOfService: DateOfService, DiagnosCode: DiagnosCode, ProcedureCode: ProcedureCode,	TypeOfService: TypeOfService, ClaimStatus: "Requested"}
	
	LCBytes, err := json.Marshal(LC)
    APIstub.PutState(ClaimID,LCBytes)

	fmt.Println("Claim Requested -> ", LC)

	return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) approveClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	LCAsBytes, _ := APIstub.GetState(ClaimID)

	var lc InsuranceClaim

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Claim json unmarshaling")
	}


	//LC := SmartContract{ClaimID: lc.ClaimID, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Issued"}
	LC := InsuranceClaim{ClaimID: lc.ClaimID, PolicyNumber: lc.PolicyNumber, EntryDate: lc.EntryDate, InsuranceCompany: lc.InsuranceCompany, PlaceOfService: lc.PlaceOfService,	ProviderName: lc.ProviderName, ClaimAmount: lc.ClaimAmount, DateOfService: lc.DateOfService, DiagnosCode: lc.DiagnosCode, ProcedureCode: lc.ProcedureCode,	TypeOfService: lc.TypeOfService, ClaimStatus: "Approved"}

	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Claim json marshaling")
	}

    APIstub.PutState(lc.ClaimID,LCBytes)
	fmt.Println("Claim Approved -> ", LC)


	return shim.Success(nil)
}

func (s *SmartContract) rejectClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	

	LCAsBytes, _ := APIstub.GetState(ClaimID)

	var lc InsuranceClaim

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Claim json unmarshaling")
	}


	//LC := SmartContract{ClaimID: lc.ClaimID, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
	LC := InsuranceClaim{ClaimID: lc.ClaimID, PolicyNumber: lc.PolicyNumber, EntryDate: lc.EntryDate, InsuranceCompany: lc.InsuranceCompany, PlaceOfService: lc.PlaceOfService,	ProviderName: lc.ProviderName, ClaimAmount: lc.ClaimAmount, DateOfService: lc.DateOfService, DiagnosCode: lc.DiagnosCode, ProcedureCode: lc.ProcedureCode,	TypeOfService: lc.TypeOfService, ClaimStatus: "Rejected"}


	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Claim json marshaling")
	}

    APIstub.PutState(lc.ClaimID,LCBytes)
	fmt.Println("Claim Rejected -> ", LC)


	

	return shim.Success(nil)
}

func (s *SmartContract) adjudicateClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	

	LCAsBytes, _ := APIstub.GetState(ClaimID)

	var lc InsuranceClaim

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Claim json unmarshaling")
	}


	//LC := SmartContract{ClaimID: lc.ClaimID, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
	LC := InsuranceClaim{ClaimID: lc.ClaimID, PolicyNumber: lc.PolicyNumber, EntryDate: lc.EntryDate, InsuranceCompany: lc.InsuranceCompany, PlaceOfService: lc.PlaceOfService,	ProviderName: lc.ProviderName, ClaimAmount: lc.ClaimAmount, DateOfService: lc.DateOfService, DiagnosCode: lc.DiagnosCode, ProcedureCode: lc.ProcedureCode,	TypeOfService: lc.TypeOfService, ClaimStatus: "Processing"}


	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Claim json marshaling")
	}

    APIstub.PutState(lc.ClaimID,LCBytes)
	fmt.Println("Claim Processing -> ", LC)


	

	return shim.Success(nil)
}

func (s *SmartContract) acceptClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	

	LCAsBytes, _ := APIstub.GetState(ClaimID)

	var lc InsuranceClaim

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with Claim json unmarshaling")
	}


	//LC := SmartContract{ClaimID: lc.ClaimID, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
	LC := InsuranceClaim{ClaimID: lc.ClaimID, PolicyNumber: lc.PolicyNumber, EntryDate: lc.EntryDate, InsuranceCompany: lc.InsuranceCompany, PlaceOfService: lc.PlaceOfService,	ProviderName: lc.ProviderName, ClaimAmount: lc.ClaimAmount, DateOfService: lc.DateOfService, DiagnosCode: lc.DiagnosCode, ProcedureCode: lc.ProcedureCode,	TypeOfService: lc.TypeOfService, ClaimStatus: "Accepted"}


	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with Claim json marshaling")
	}

    APIstub.PutState(lc.ClaimID,LCBytes)
	fmt.Println("Claim Accepted -> ", LC)


	

	return shim.Success(nil)
}


func (s *SmartContract) getClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(ClaimID)

	return shim.Success(LCAsBytes)
}

func (s *SmartContract) getClaimHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	ClaimID := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(ClaimID)
	if err != nil {
		return shim.Error("Error retrieving Claim history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving Claim history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getClaimHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
