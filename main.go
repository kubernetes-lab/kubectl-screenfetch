package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gilliek/go-xterm256/xterm256"

	"github.com/kubernetes-lab/kubectl-screenfetch/pkg/k8s"
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
	logo := `                                                  
                     ,:i11i:,                     
                 ,:ittffttfftti:,                 
            .,;ittffttttttttttfftti;,.            
        .,;1tffftttttttt00ttttttttffft1;,.        
      :1tfffttttttttttt1881tttttttttttffft1:      
     :ftttttttttttttfLLC@8CLfftttttttttttttf:     
     1tttttG0LtttLG8@@@@@@@@@@8GLtttL0Gttttt1     
    ,ftttttLG80G0@@0CLtL@@LtLC0@@0G08GLtttttf,    
    1ttttttt1t8@@@Gf111C@@C111fG@@@8t1ttttttt1    
   ,ftttttttt0@8C0@@8Cf0@@0fC8@@0C8@0ttttttttf,   
   iftttttttL@@L1tfG@@@@@@@@@@Gft1L@@Ltttttttfi   
  .tttttttttG@8tttff0@@GffG@@0ffttt8@0ttttttttt.  
  iftttttttt0@@08@@@@@@0LL0@@@@@@80@@0ttttttttfi  
 .tttttf08008@@@GGCLf0@@@@@@0fLCCG@@@80080fttttt. 
 ,ftttttLLfttG@@L111f8@@LL@@8f111L@@GttfLLtttttf, 
  :tftttttttttC@@0Lf8@8f11f8@8fL0@@Ctttttttttft:  
   .ittttttttttfG@@@@@CffffC@@@@@Gftttttttttti.   
     ,1ftttttttttt0@8@@@@@@@@8@0ttttttttttf1,     
       ;tftttttttL@GtffLLLLfftG@Ctttttttft;       
        .iftttttf80tttttttttttt08ftttttfi.        
          :tfttttttttttttttttttttttttft:          
           .itfttttttttttttttttttttfti.           
             ,i11111111111111111111i,             
                                                  `
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
	ks := k8s.GetKubernetesInfo(context.Background())

	for _, infoMap := range informationList {
		for k, v := range infoMap {
			switch k {
			case K_CONTEXT:
				s = append(s, xterm256.Sprint(userColor, ks.CurrentContext))
			case K_SEP:
				s = append(s, xterm256.Sprint(sepColor, "--------------------------------"))
			case K_VERSION:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.Version))
			case K_NODES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.NodesCount))
			case K_NAMESPACES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.NamespacesCount))
			case K_DEPLOYMENTS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.DeploymentsCount))
			case K_PODS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.PodsCount))
			case K_SERVICES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.ServicesCount))
			case K_INGRESSES:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.IngressesCount))
			case K_PVS:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.PVsCount))
			case K_CRI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.CRI))
			case K_CNI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.CNI))
			case K_CSI:
				s = append(s, xterm256.Sprint(titleColor, v+": ")+xterm256.Sprint(infoColor, ks.CSI))
			}
		}
	}
	return s
}
