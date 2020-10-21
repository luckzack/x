package sha256

import (
	"fmt"
	"testing"
	"log"
	"time"
)

func Test_enc(t *testing.T){

	ts := time.Now().Unix() + 10
	s := fmt.Sprintf("4a62eb464a508d42hello32bce3a2-9e02-4518-9145-8ab6df853de5%dnaKaDk78tZvnDc741hDFPOwiFwSsvJ6W", ts)

	log.Println(ts, EncryptString(s))
}