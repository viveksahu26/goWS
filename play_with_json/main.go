package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-openapi/strfmt"
	"github.com/sigstore/rekor/pkg/generated/models"
)

type RekorBundle struct {
	SignedEntryTimestamp []byte
	Payload              RekorPayload
}

type RekorPayload struct {
	Body           interface{} `json:"body"`
	IntegratedTime int64       `json:"integratedTime"`
	LogIndex       int64       `json:"logIndex"`
	LogID          string      `json:"logID"`
}

type LocalSignedPayload struct {
	Base64Signature string       `json:"base64Signature"`
	Cert            string       `json:"cert,omitempty"`
	Bundle          *RekorBundle `json:"rekorBundle,omitempty"`
}
type RekorResponse map[string]LogEntryAnon

// pkg/generated/models/log_entry.go
type LogEntryAnon struct {
	// attestation
	Attestation *LogEntryAnonAttestation `json:"attestation,omitempty"`

	// body
	// Required: true
	Body interface{} `json:"body"`

	// The time the entry was added to the log as a Unix timestamp in seconds
	// Required: true
	IntegratedTime *int64 `json:"integratedTime"`

	// This is the SHA256 hash of the DER-encoded public key for the log at the time the entry was included in the log
	// Required: true
	// Pattern: ^[0-9a-fA-F]{64}$
	LogID *string `json:"logID"`

	// log index
	// Required: true
	// Minimum: 0
	LogIndex *int64 `json:"logIndex"`

	// verification
	Verification *LogEntryAnonVerification `json:"verification,omitempty"`
}

type LogEntryAnonAttestation struct {
	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`
}

type LogEntryAnonVerification struct {
	// inclusion proof
	InclusionProof *InclusionProof `json:"inclusionProof,omitempty"`

	// Signature over the logID, logIndex, body and integratedTime.
	// Format: byte
	SignedEntryTimestamp strfmt.Base64 `json:"signedEntryTimestamp,omitempty"`
}
type InclusionProof struct {
	// The checkpoint (signed tree head) that the inclusion proof is based on
	// Required: true
	Checkpoint *string `json:"checkpoint"`

	// A list of hashes required to compute the inclusion proof, sorted in order from leaf to root
	// Required: true
	Hashes []string `json:"hashes"`

	// The index of the entry in the transparency log
	// Required: true
	// Minimum: 0
	LogIndex *int64 `json:"logIndex"`

	// The hash value stored at the root of the merkle tree at the time the proof was generated
	// Required: true
	// Pattern: ^[0-9a-fA-F]{64}$
	RootHash *string `json:"rootHash"`

	// The size of the merkle tree at the time the inclusion proof was generated
	// Required: true
	// Minimum: 1
	TreeSize *int64 `json:"treeSize"`
}
type LivingBeing struct {
	Person *Person
	Dog    string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rekorResponseRef := "rekor-response.json"
	rekorResponseByte, err := os.ReadFile(filepath.Clean(rekorResponseRef))
	if err != nil {
		fmt.Println("err: ", err)
	}
	// var rekorBundle *RekorBundle
	// rekorBundle := &RekorBundle{}
	rekorBundle := new(RekorBundle)
	// var localCosignPayload LocalSignedPayload

	rekorResponse := make(models.LogEntry)
	logIndexValue := int64(63771522)
	rootHash := "5bbff4e9f8034a33b102996271c0e01b87caa83c44c46e893b47b81467fd808c"
	treeSize := int64(64319120)
	checkPoint := "rekor.sigstore.dev - 2605736670972794746\n64319120\nW7/06fgDSjOxApliccDgG4fKqDxExG6JO0e4FGf9gIw=\nTimestamp: 1706845443714164712\n\n— rekor.sigstore.dev wNI9ajBFAiEAnduIhP1Jjz8E0ZAP8e1x0aKqzJCtmWZyV1mRJB/PlOoCIBoeHdjeONYmxlD2Za7sU0NeK/60skNnwoelsa3m2M8z\n"
	integratedTime := int64(1706680021)
	logID := "c0d23d6ad406973f9559f3ba2d1ca01f84147d8ffc5b8445c224f98b9591801d"
	logIndex := int64(67934953)

	logEntry := models.LogEntryAnon{
		Body:           "eyJhcGlWZXJzaW9uIjoiMC4wLjEiLCJraW5kIjoiaGFzaGVkcmVrb3JkIiwic3BlYyI6eyJkYXRhIjp7Imhhc2giOnsiYWxnb3JpdGhtIjoic2hhMjU2IiwidmFsdWUiOiI3YTIxODlmNzNlYTVkYmVlMmQ1NjgxMGU4NjBjZDgxNjdlYTFiMGYzMTdkZGRjMmU0YzE2NmU3ZDY4NzUzOGYwIn19LCJzaWduYXR1cmUiOnsiY29udGVudCI6Ik1FUUNJQlZtMVc1YlNSOXBTYVFyRWRVL2dOeEZuaVQzMCs4RGp5SG9naERwWHM3b0FpQXMyby82c3l2bzRaTDd3YVUrMXBNeVIvR1dNWlZTaUdzRm5hSm04ZDZZZ0E9PSIsInB1YmxpY0tleSI6eyJjb250ZW50IjoiTFMwdExTMUNSVWRKVGlCUVZVSk1TVU1nUzBWWkxTMHRMUzBLVFVacmQwVjNXVWhMYjFwSmVtb3dRMEZSV1VsTGIxcEplbW93UkVGUlkwUlJaMEZGUldGSWNHUmhZVE13ZEVOaVMxbG5lalpyUlhsQmNqTXJORmh5TWdwb1pWUjNaM2hQVVVrM2QwcDRiVWhIWW5OU1dYcEJWbFJyUVN0RGIzVjZUblZrTUc5eGRuUTRXVXB0Tm5CQlFUZG5SbUZFUkZGU05EUlJQVDBLTFMwdExTMUZUa1FnVUZWQ1RFbERJRXRGV1MwdExTMHRDZz09In19fX0=",
		IntegratedTime: &integratedTime,
		LogID:          &logID,
		LogIndex:       &logIndex,
		Verification: &models.LogEntryAnonVerification{
			InclusionProof: &models.InclusionProof{
				Checkpoint: &checkPoint,
				LogIndex:   &logIndexValue,
				RootHash:   &rootHash,
				TreeSize:   &treeSize,
			},
			SignedEntryTimestamp: strfmt.Base64("MEUCIQCLiNiSLxFk8vgkCopYcuFQXGEcvr6YM0TXgFUe5HcAHAIgSJuWj3uH8QrQeaEfc5ddMpIwU4JdXmQD3bgfkntEVTk="),
		},
	}
	rekorResponse["24296fb24b8ad77a8ada322cbba201d23b88acd2f68b29e358668e3aef36306ddb2c253ca0dc9ede"] = logEntry
	jsonRekorResponsePath, err := json.Marshal(rekorResponse)
	if err != nil {
		fmt.Println("Error While Marshaling rekorResponse: ", err)
	}

	rekorResponsePath := filepath.Join("/home/linuzz/go/src/github.com/viveksahu26/goWs/play_with_json", "test-rekor-response.json")
	if err := os.WriteFile(rekorResponsePath, jsonRekorResponsePath, 0o644); err != nil {
		fmt.Println("ERROR while writting to a file: ", err)
	}
	rekorResponseByteTest, err := os.ReadFile(filepath.Clean(rekorResponsePath))
	if err != nil {
		fmt.Println("err: ", err)
	}

	var rekorResponsesTest RekorResponse
	fmt.Println("rekorResponseByteTest: ", string(rekorResponseByteTest))

	err = json.Unmarshal(rekorResponseByteTest, &rekorResponsesTest)
	if err != nil {
		fmt.Println("Unmarshal rekorResponsesTest error ", err)
	}
	// non-nil
	fmt.Println("rekorResponsesTest: ", rekorResponsesTest)
	for key, v := range rekorResponsesTest {
		fmt.Println("Key: \n", key)

		// nil
		if v.Attestation == nil {
			fmt.Println("rekorResponsesTest.Attestation is nil ")
		} else {
			fmt.Println("rekorResponsesTest.Attestation: ", v.Attestation)
		}

		// non-nil
		if v.Body == nil {
			fmt.Println("v.Body is nil ")
		} else {
			fmt.Println("v.Body: ", v.Body)
			rekorBundle.Payload.Body = v.Body
		}

		// non-nil
		if v.IntegratedTime == nil {
			fmt.Println("v.Integrated is nil ")
		} else {
			fmt.Println("v.IntegratedTime: ", *v.IntegratedTime)
			rekorBundle.Payload.IntegratedTime = *v.IntegratedTime
		}

		// non-nil
		if v.LogID == nil {
			fmt.Println("v.LogID is nil ")
		} else {
			fmt.Println("v.LogID: ", *v.LogID)
			rekorBundle.Payload.LogID = *v.LogID
		}

		// non-nil
		if v.LogIndex == nil {
			fmt.Println("v.LogIndex is nil ")
		} else {
			fmt.Println("v.LogIndex: ", *v.LogIndex)
			rekorBundle.Payload.LogIndex = *v.LogIndex
		}

		// non-nil
		if v.Verification.SignedEntryTimestamp == nil {
			fmt.Println("v.Verification.SignedEntryTimestamp is nil ")
		} else {
			fmt.Println("v.Verification.SignedEntryTimestamp: ", v.Verification.SignedEntryTimestamp)
			rekorBundle.SignedEntryTimestamp = v.Verification.SignedEntryTimestamp
		}
		// non-nil
		if v.Verification.InclusionProof == nil {
			fmt.Println("v.Verification.InclusionProof is nil ")
		} else {
			fmt.Println("v.Verification.InclusionProof: ", *v.Verification.InclusionProof)
		}

		// non-nil
		if v.Verification == nil {
			fmt.Println("v.Verification is nil ")
		} else {
			fmt.Println("v.Verification: ", *v.Verification)
		}

		// rekorBundle.SignedEntryTimestamp = v.Verification.SignedEntryTimestamp
		// rekorBundle.Payload.Body = v.Body
		// rekorBundle.Payload.IntegratedTime = *v.IntegratedTime
		// rekorBundle.Payload.LogIndex = *v.LogIndex
		// rekorBundle.Payload.LogID = *v.LogID
		fmt.Println("rekorBundle: ", rekorBundle)

	}

	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------------------------------------")

	var rekorResponses RekorResponse
	fmt.Println("rekorResponseByte: ", string(rekorResponseByte))

	err = json.Unmarshal(rekorResponseByte, &rekorResponses)
	if err != nil {
		fmt.Println("Unmarshal rekorResponse error ", err)
	}
	// non-nil
	fmt.Println("rekorResponse: ", rekorResponse)
	for key, v := range rekorResponse {
		fmt.Println("Key: \n", key)

		// nil
		if v.Attestation == nil {
			fmt.Println("rekorResponse.v is nil ")
		} else {
			fmt.Println("rekorResponse.v: ", v.Attestation)
		}

		// non-nil
		if v.Body == nil {
			fmt.Println("v.Body is nil ")
		} else {
			fmt.Println("v.Body: ", v.Body)
		}

		// non-nil
		if v.IntegratedTime == nil {
			fmt.Println("v.Integrated is nil ")
		} else {
			fmt.Println("v.IntegratedTime: ", v.IntegratedTime)
		}

		// non-nil
		if v.LogID == nil {
			fmt.Println("v.LogID is nil ")
		} else {
			fmt.Println("v.LogID: ", v.LogID)
		}

		// non-nil
		if v.LogIndex == nil {
			fmt.Println("v.LogIndex is nil ")
		} else {
			fmt.Println("v.LogIndex: ", v.LogIndex)
		}

		// non-nil
		if v.Verification.SignedEntryTimestamp == nil {
			fmt.Println("v.Verification.SignedEntryTimestamp is nil ")
		} else {
			fmt.Println("v.Verification.SignedEntryTimestamp: ", v.Verification.SignedEntryTimestamp)
		}
		// non-nil
		if v.Verification.InclusionProof == nil {
			fmt.Println("v.Verification.InclusionProof is nil ")
		} else {
			fmt.Println("v.Verification.InclusionProof: ", v.Verification.InclusionProof)
		}

		// non-nil
		if v.Verification == nil {
			fmt.Println("v.Verification is nil ")
		} else {
			fmt.Println("v.Verification: ", v.Verification)
		}
	}
	// if err := json.Unmarshal(rekorResponseByte, &localCosignPayload); err == nil {
	// 	rekorBundle = localCosignPayload.Bundle
	// 	if rekorBundle == nil && localCosignPayload.Cert == "" && localCosignPayload.Base64Signature == "" {
	// 		fmt.Println("rekor-bundle is passed")
	// err = json.Unmarshal(rekorResponseByte, &rekorBundle)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// 		} else {
	// 			fmt.Println("NO MARSHAL ERROR")
	// 		}
	// 	} else {
	// 		fmt.Println("Bundle is passed")
	// 	}
	// } else {
	// 	fmt.Println("ELSE")
	// 	fmt.Printf("error: %v\n", err)
	// }

	// if rekorBundle == nil {
	// 	fmt.Println("error: unable to parse Rekor response to attach to image")
	// } else {
	// 	fmt.Println("rekorBundle is provided by the user but it's value is nil")
	// }
	// fmt.Println("rekorBundle: ", rekorBundle)
}
