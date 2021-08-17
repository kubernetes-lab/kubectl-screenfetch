package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gilliek/go-xterm256/xterm256"
)

const (
	K_CONTEXT     = "context"
	K_SEP         = "sep"
	K_VERSION     = "version"
	K_NODES       = "nodes"
	K_NAMESPACES  = "namespaces"
	K_DEPLOYMENTS = "deployments"
	K_PODS        = "pods"
	K_SERVICES    = "services"
	K_INGRESSES   = "ingresses"
	K_PVS         = "pvs"
	K_CRI         = "cri"
	K_CNI         = "cni"
	K_CSI         = "csi"
)

var informationList []map[string]string

func init() {
	// Init "Info: InfoTitle"
	informationList = []map[string]string{
		{K_CONTEXT: ""},
		{K_SEP: ""},
		{K_VERSION: "Kubernetes Version"},
		{K_NODES: "Nodes"},
		{K_NAMESPACES: "Namespaces"},
		{K_DEPLOYMENTS: "Deployments"},
		{K_PODS: "Pods"},
		{K_SERVICES: "Services"},
		{K_INGRESSES: "Ingresses"},
		{K_PVS: "PVs"},
		{K_CRI: "CRI"},
		{K_CNI: "CNI"},
		{K_CSI: "CSI"},
	}
}

func main() {
	logo := `                        ..                        
                   .:;i1tt1i;:.                   
             .,:;1ttfffttttffftt1;:,.             
        .,:i1ttfffttttttGGttttttffftt1i:,.        
      ;1tffttttttttttttt88tttttttttttttfft1;      
     ;fttttLLtttttfLCG00@@00GCLftttttLLttttf;     
    .ttttttC00CLC08800CG@@GC00880CLC00Ctttttt.    
    iftttttttf@@@@0ft11L@@L11tf0@@@@ftttttttfi    
   ,ftttttttt0@8LG@@8GL0@@0LG8@@GL8@0ttttttttf,   
   1ttttttttC@@f11tL8@@@00@@@8Lt11f@@Ctttttttt1   
  :ftttttttt0@8CG008@@@GttG@@@800GC8@0ttttttttf:  
 .ttttttG0GG8@@@0GGGL0@@@@@@0LGGG0@@@8GG0Gtttttt. 
 ,tfttttLLfftC@@Ct11f8@8LL8@8f11tC@@CtffLLttttft, 
  .itfttttttttf0@@GC@@Gt11tG@@CG@@0fttttttttfti.  
    .itfttttttttfL8@@@800000@@@8Cfttttttttfti.    
      .;tftttttttL8GfLCCCCCCLfG8Ltttttttfti.      
        .;tfttttf0GttttttttttttG0fttttft;.        
          .;tfttttttttttttttttttttttft;.          
            .;1tttttttttttttttttttt1;.            
               ....................    `
	logoSlice := strings.Split(logo, "\n")

	s := getInformation(xterm256.Green, xterm256.White, xterm256.Red, xterm256.Red)
	scanner := bufio.NewScanner(strings.NewReader(""))
	scanner = bufio.NewScanner(strings.NewReader(logo))

	index := 0
	for i, str := range s {
		if len(logoSlice)-1 < i {
			fmt.Println(strings.Repeat(" ", utf8.RuneCountInString(logoSlice[0])) + "    " + str)
		} else {
			fmt.Println(xterm256.Sprint(xterm256.Blue, logoSlice[i]) + "    " + str)
		}
	}
	for scanner.Scan() {
		if index >= len(s) {
			fmt.Println(xterm256.Sprint(xterm256.Blue, scanner.Text()))
		}
		index++
	}
}

func getInformation(titleColor xterm256.Color, infoColor xterm256.Color, userColor xterm256.Color, sepColor xterm256.Color) (s []string) {
	for _, infoMap := range informationList {
		for k, v := range infoMap {
			switch k {
			case K_CONTEXT:
				s = append(s, xterm256.Sprint(userColor, "user.Username"))
			case K_SEP:
				s = append(s, xterm256.Sprint(sepColor, "--------------------------------"))
			case K_VERSION:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, "v1.21.1"))
			case K_NODES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_NAMESPACES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_DEPLOYMENTS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_PODS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_SERVICES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_INGRESSES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_PVS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_CRI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_CNI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			case K_CSI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, 7))
			}
		}
	}
	return s
}
