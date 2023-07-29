package __commonpkg

import (
	"encoding/json"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	v := `{"code":"SLB.0000","message":"success","data":{"loadbalancers":[{"name":"lb-ceake","loadbalancerId":"lb-jsc25aeg3x9a","status":"Running","ipv4Address":"192.168.0.5","regionId":"region0","zoneIds":["region0"],"description":"","createTime":"2023-03-08 14:37:36","updateTime":"2023-03-08 14:38:16","vpcId":"vpc-cxcegpe7r0lvagpqjlxt","subnetId":"snet-hrt9ebdaqa5r8w5a5fwl","eip":null,"expireTime":"2023-03-08 14:37:36","payModel":"","orderId":"","spec":"slb.x1","ipv6":null,"listeners":[{"listenerId":"ls-kzii0jdcts42","port":6443,"protocol":"TCP"}],"acrossVpc":[{"ipv4Address":"192.168.0.6","Ipv6Address":""},{"ipv4Address":"192.168.0.7","Ipv6Address":""}]}],"total":1,"requestId":"f3e770ed-b73e-4411-a851-a09ee0a23670"}}`
	res := map[string]interface{}{}
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		t.Fatal(err)
	}
}
