/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Kurl <URL>",
	Short: "curl cli ",
	Long: `A curl cli build for learning can be used to make HTTP request to a server`,

	Args:cobra.ExactArgs(1),

    Run: func(cmd *cobra.Command, args []string) { 
		u,err:=url.Parse(args[0])
		if err!=nil{
			panic(err)
		}
		host:=u.Hostname()
		path:=u.Path
		port:=u.Port()

		if port==""{
			port="80"
		}
     
		con,err:=net.Dial("tcp",fmt.Sprintf("%s:%s",host,port))
		if err!=nil{
			panic(err)
		}

	    defer con.Close()
		
	    fmt.Fprintf(con, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", path, host)
		
		buf:=make([]byte,1024)
		n,err:=con.Read(buf)
		if err!=nil{
			panic(err)
		}
		        
		 fmt.Println(string(buf[:n]))

	},
}

func Execute(){
	err:=rootCmd.Execute()
	if err!=nil{
		os.Exit(1)
	}
}
