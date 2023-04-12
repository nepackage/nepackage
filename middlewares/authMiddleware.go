package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/auth"
	"github.com/nepackage/nepackage/repository"
	"github.com/nepackage/nepackage/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		/* Validate that a JWT exists in Authorization header, if not return a 401 error message */
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		/* Get claims in JWT, if found an error return a 401 error message */
		jwtClaims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		/* Get roles in JWT provided by user in Authorization header */
		userRoles := utils.ConvertStringToUintStruct(jwtClaims.Role)
		pathRequested := c.Request.URL.Path
		methodRequested := c.Request.Method
		rolePolicies, err := repository.GetRolePolicies(userRoles)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		/* Get role's policies with path requested */
		policiesWithPathRequested, err := utils.GetPoliciesWithMatchedPath(rolePolicies, pathRequested)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		/* Compare path requested with path in AuthorizedRoles */
		authorized, err := utils.ValidateMethodRequestWithPolicyMethod(policiesWithPathRequested, methodRequested)
		log.Println(authorized)
		log.Println("========================================")

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You're not able to perform " + methodRequested + " method on: " + pathRequested})
			c.Abort()
			return
		}

		c.Next()
	}
}
