package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"net/http"
)

var rootCmd = &cobra.Command{
	Use:   "spas",
	Short: "Signal page application server",
	Long:  `Signal page application testing server`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.LoadHTMLGlob("./index.html")
		r.Use(static.Serve("/", static.LocalFile(".", true)))
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		r.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
