package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MemberInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type MemberClient struct {
	BaseURL string
}

func NewMemberClient(url string) *MemberClient {
	return &MemberClient{BaseURL: url}
}

func (mc *MemberClient) GetMember(memberID uint) (*MemberInfo, error) {
	url := fmt.Sprintf("%s/members/%d", mc.BaseURL, memberID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("member not found")
	}

	var member MemberInfo
	if err := json.NewDecoder(resp.Body).Decode(&member); err != nil {
		return nil, err
	}

	return &member, nil
}
