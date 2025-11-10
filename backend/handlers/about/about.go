package about
import (
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetAboutMD(c *gin.Context) {
    content, err := ioutil.ReadFile("/app/static/about/about1.md")
    if err != nil {
        c.String(http.StatusInternalServerError, "无法读取文件")
        return
    }
    c.String(http.StatusOK, string(content))
}
