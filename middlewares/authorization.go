package middlewares

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Authorize returns a middleware that checks whether the authenticated user has the given permission.
// Usage: router.GET("/foo", middlewares.Authenticate(jwtService), middlewares.Authorize(rbacService, "perm_name"), handler)
func Authorize(rbacSvc service.RbacService, permission string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        uid, exists := ctx.Get("user_id")
        if !exists {
            res := utils.BuildResponseFailed("authorization failed", "user not found in context", nil)
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
            return
        }

        userIDStr, ok := uid.(string)
        if !ok {
            res := utils.BuildResponseFailed("authorization failed", "invalid user id", nil)
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
            return
        }

        userUUID, err := uuid.Parse(userIDStr)
        if err != nil {
            res := utils.BuildResponseFailed("authorization failed", "invalid user id format", nil)
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
            return
        }

        // retrieve roles with permissions
        roles, err := rbacSvc.GetRolesByUser(ctx.Request.Context(), nil, userUUID)
        if err != nil {
            res := utils.BuildResponseFailed("authorization failed", err.Error(), nil)
            ctx.AbortWithStatusJSON(http.StatusForbidden, res)
            return
        }

        // check permission
        allowed := false
        for _, r := range roles {
            for _, p := range r.Permissions {
                if p.Name == permission {
                    allowed = true
                    break
                }
            }
            if allowed {
                break
            }
        }

        if !allowed {
            res := utils.BuildResponseFailed("authorization failed", "permission denied", nil)
            ctx.AbortWithStatusJSON(http.StatusForbidden, res)
            return
        }

        ctx.Next()
    }
}
