package main

import (
  			"fmt"
				"time"
				"log"
				"net"
				//"runtime"
				//"strings"
				//"os"
				//"regexp"
)

func time_test(){
	t :=time.Now()
	fmt.Println(t)
	fmt.Println(t)
	fmt.Println("%d\n",t.Year())
}

/*
func isTimeout(err error) bool {
	e, ok := err.(error)
	return ok && e.Timeout()
}
*/

func ftpcmd(cmd string,result []uint8,sock net.Conn){
    
    //send cmd
    cmd +="\r\n"
    send_buf :=[]byte(cmd)
    buf_len :=0
    
    //fmt.Println("%s\n",send_buf)
    sock.SetDeadline(time.Now().Add(100 * time.Millisecond))
    send_len,err :=sock.Write(send_buf)
    if err != nil {
		log.Fatal(err)
	}
	
    log.Printf("send:%d,%s",send_len,send_buf[:])
    
    //loop read the result of cmd
    for{
      sock.SetDeadline(time.Now().Add(1000 * time.Millisecond))
      read_len,err :=sock.Read(result[buf_len:])
	    if err != nil {
	    /*
	        if err.Timeout(){
              log.Printf("RECV time out")
      	      break
            }
        */
        
            //log.Printf("RECV time out")
//		    log.Fatal(err)
		    break
	    }
	    if(read_len==0){
	        break
	    }
	   //log.Printf("RECV:read_len %d,%s",read_len,result[buf_len:read_len])
	    buf_len+=read_len
		}
		//log.Printf("RECV:%d,%q",read_len,result[:])
		log.Printf("RECV:%d,%s",buf_len,result[:buf_len])
}



var ftp_conn net.Conn
/*
var user_name = "admin"
var user_passwd = "zhongxing"
var	ftp_addr ="192.168.1.5:21"
*/
/*
var user_name = "anonymous"
var user_passwd = "kevin@gmail.com"
var	ftp_addr ="ftp.freebsd.org:21"
*/

var user_name = "anonymous"
var user_passwd = "kevin@gmail.com"
var	ftp_addr ="127.0.0.1:21"

func main(){


	data := make([]uint8, 4096)
	read_buf := make([]uint8, 4096) 

	ftp_conn, err := net.Dial("tcp", ftp_addr)
	if err != nil {
		log.Fatal(err)
	}
	
	
	defer 	ftp_conn.Close()
	
	log.Println("Connect Server: " + ftp_addr + " OK\n")
	
	/* need loop for read sock buf */
	read_len,err :=ftp_conn.Read(data[:])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server Response:%d,%s",read_len,data[:read_len])
	
	ftpcmd("USER "+user_name,read_buf,ftp_conn)
	
    ftpcmd("PASS "+user_passwd,read_buf,ftp_conn)
    
    //ftpcmd("PORT ",read_buf,ftp_conn)
    
    ftpcmd("SYST ",read_buf,ftp_conn)
    
    //ftpcmd("TYPE A ",read_buf,ftp_conn)
    
    //ftpcmd("PASV ",read_buf,ftp_conn)

    
    ftpcmd("LIST /",read_buf,ftp_conn)
    ftpcmd("PWD ",read_buf,ftp_conn)
    
	ftpcmd("QUIT ",read_buf,ftp_conn)
}
