/**
 * 获取权限配置
 * 
 * @author xiaoRui
*/
import { checkAuthority } from "@/utils/authority"
import storage from "@/utils/storage"

export default {
    inserted(el, binding, vnode) {
        const { value } = binding
        const permissions = storage.getItem('permissionList') || []
        const hasPermission = checkAuthority(value, permissions)
        if (!hasPermission) {
            if (el.parentNode) {
                el.parentNode.removeChild(el)
            } else {
                el.innerHTML = ''
            }
        } else {
            el && el.setAttribute('code', value)
        }
    }
}