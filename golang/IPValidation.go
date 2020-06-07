package kata

import("regexp")
import("strings")
import("strconv")

func Is_valid_ip(ip string) bool {
  matched, _ := regexp.MatchString(`(0|([1-9](\d){0,2}))\.(0|([1-9](\d){0,2}))\.(0|([1-9](\d){0,2}))\.(0|([1-9](0[1-9]|[1-9]{0,2}){0,1}))$`, ip)
  
  for _, v := range strings.Split(ip, "."){
    val, _:= strconv.Atoi(v)
    if(val >255){
      matched=false
    }
  }
  return matched
}

/*
Write an algorithm that will identify 
valid IPv4 addresses in dot-decimal format. 
IPs should be considered valid if 
they consist of four octets, 
with values between 0 and 255, inclusive.
*/