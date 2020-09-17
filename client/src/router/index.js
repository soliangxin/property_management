import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },

  /* 小区管理
       --> 查看小区
       --> 新增小区
  */
  {
    path: '/community',
    component: Layout,
    redirect: '/community/ViewCommunity',
    name: 'Community',
    meta: { title: '小区管理', icon: 'guide' },
    children: [
      {
        path: 'viewCommunity',
        name: 'ViewCommunity',
        component: () => import('@/views/community/view_community'),
        meta: { title: '查看小区', icon: 'component' }
      },
      {
        path: 'addCommunity',
        name: 'AddCommunity',
        component: () => import('@/views/community/add_community'),
        meta: { title: '新增小区', icon: 'edit' }
      }
    ]
  },

  /* 房产信息管理
      --> 查看房产
      --> 新增房产
 */
  {
    path: '/house',
    component: Layout,
    redirect: '/house/ViewHouse',
    name: 'House',
    meta: { title: '房产信息管理', icon: 'list' },
    children: [
      {
        path: 'viewHouse',
        name: 'ViewHouse',
        component: () => import('@/views/house/view_house'),
        meta: { title: '查看房产', icon: 'component' }
      },
      {
        path: 'addHouse',
        name: 'AddHouse',
        component: () => import('@/views/house/add_house'),
        meta: { title: '新增业主', icon: 'edit' }
      }
    ]
  },

  /* 住户信息管理
     --> 查看住户信息
     --> 新增住户信息
*/
  {
    path: '/tenement',
    component: Layout,
    redirect: '/tenement/AddTenement',
    name: 'Tenement',
    meta: { title: '住户信息管理', icon: 'user' },
    children: [
      {
        path: 'viewTenement',
        name: 'ViewTenement',
        component: () => import('@/views/tenement/view_tenement'),
        meta: { title: '查看住户', icon: 'component' }
      },
      {
        path: 'addTenement',
        name: 'AddTenement',
        component: () => import('@/views/tenement/add_tenement'),
        meta: { title: '新增住户', icon: 'edit' }
      }
    ]
  },

  /* 停车位管理
      --> 查看车位
      --> 新增车位
  */
  {
    path: '/stall',
    component: Layout,
    redirect: '/stall/AddStall',
    name: 'Stall',
    meta: { title: '停车位管理', icon: 'lock' },
    children: [
      {
        path: 'viewStall',
        name: 'ViewStall',
        component: () => import('@/views/stall/view_stall'),
        meta: { title: '停车位使用情况', icon: 'component' }
      },
      {
        path: 'addStall',
        name: 'AddStall',
        component: () => import('@/views/stall/add_stall'),
        meta: { title: '新增停车位', icon: 'edit' }
      }
    ]
  },

  /* 收费管理
       --> 缴费记录管理
       --> 收费项目
       --> 新增收费项目
  */
  {
    path: '/charge',
    component: Layout,
    redirect: '/charge/viewRecordCharge',
    name: 'Charge',
    meta: { title: '收费管理', icon: 'money' },
    children: [
      {
        path: 'viewRecordCharge',
        name: 'ViewRecordCharge',
        component: () => import('@/views/charge/view_record_charge'),
        meta: { title: '缴费记录管理', icon: 'chart' }
      },
      {
        path: 'userCharge',
        name: 'UserCharge',
        component: () => import('@/views/charge/user_pay_fee'),
        meta: { title: '用户缴费', icon: 'skill' }
      },
      {
        path: 'viewCharge',
        name: 'ViewCharge',
        component: () => import('@/views/charge/view_item_charge'),
        meta: { title: '收费项目', icon: 'documentation' }
      },
      {
        path: 'addCharge',
        name: 'AddCharge',
        component: () => import('@/views/charge/add_charge'),
        meta: { title: '新增收费项目', icon: 'edit' }
      }
    ]
  },

  /* 资产设备管理
     --> 资产查看
     --> 新增资产
 */
  {
    path: '/asset',
    component: Layout,
    redirect: '/asset/AddAsset',
    name: 'Asset',
    meta: { title: '资产设备管理', icon: 'list' },
    children: [
      {
        path: 'viewAsset',
        name: 'ViewAsset',
        component: () => import('@/views/asset/view_asset'),
        meta: { title: '资产查看', icon: 'component' }
      },
      {
        path: 'addAsset',
        name: 'AddAsset',
        component: () => import('@/views/asset/add_asset'),
        meta: { title: '新增资产', icon: 'edit' }
      }
    ]
  },

  /* 值班管理
   --> 值班安排
   --> 新增值班安排
  */
  {
    path: '/watch',
    component: Layout,
    redirect: '/watch/AddWatch',
    name: 'Watch',
    meta: { title: '值班管理', icon: 'nested' },
    children: [
      {
        path: 'viewWatch',
        name: 'ViewWatch',
        component: () => import('@/views/watch/view_watch'),
        meta: { title: '值班安排', icon: 'component' }
      },
      {
        path: 'addWatch',
        name: 'AddWatch',
        component: () => import('@/views/watch/add_watch'),
        meta: { title: '新增值班安排', icon: 'edit' }
      }
    ]
  },
  /* 系统配置
  --> 权限管理
  --> 新增用户
  --> 操作日志
  */
  {
    path: '/system',
    component: Layout,
    redirect: '/system/AddWatch',
    name: 'System',
    meta: { title: '系统配置', icon: 'tree-table' },
    children: [
      {
        path: 'authority',
        name: 'authority',
        meta: { title: '权限管理', icon: 'component' }
      },
      {
        path: 'addUser',
        name: 'AddUser',
        component: () => import('@/views/system/add_user'),
        meta: { title: '新增用户', icon: 'peoples' }
      },
      {
        path: 'editLog',
        name: 'EditLog',
        meta: { title: '操作日志', icon: 'edit' }
      }
    ]
  },
  /* 注释不需要的功能
  {
    path: '/example',
    component: Layout,
    redirect: '/example/table',
    name: 'Example',
    meta: { title: 'Example', icon: 'example' },
    children: [
      {
        path: 'table',
        name: 'Table',
        component: () => import('@/views/table/index'),
        meta: { title: 'Table', icon: 'table' }
      },
      {
        path: 'tree',
        name: 'Tree',
        component: () => import('@/views/tree/index'),
        meta: { title: 'Tree', icon: 'tree' }
      }
    ]
  },

  {
    path: '/form',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'Form',
        component: () => import('@/views/form/index'),
        meta: { title: 'Form', icon: 'form' }
      }
    ]
  },

  {
    path: '/nested',
    component: Layout,
    redirect: '/nested/menu1',
    name: 'Nested',
    meta: {
      title: 'Nested',
      icon: 'nested'
    },
    children: [
      {
        path: 'menu1',
        component: () => import('@/views/nested/menu1/index'), // Parent router-view
        name: 'Menu1',
        meta: { title: 'Menu1' },
        children: [
          {
            path: 'menu1-1',
            component: () => import('@/views/nested/menu1/menu1-1'),
            name: 'Menu1-1',
            meta: { title: 'Menu1-1' }
          },
          {
            path: 'menu1-2',
            component: () => import('@/views/nested/menu1/menu1-2'),
            name: 'Menu1-2',
            meta: { title: 'Menu1-2' },
            children: [
              {
                path: 'menu1-2-1',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1'),
                name: 'Menu1-2-1',
                meta: { title: 'Menu1-2-1' }
              },
              {
                path: 'menu1-2-2',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2'),
                name: 'Menu1-2-2',
                meta: { title: 'Menu1-2-2' }
              }
            ]
          },
          {
            path: 'menu1-3',
            component: () => import('@/views/nested/menu1/menu1-3'),
            name: 'Menu1-3',
            meta: { title: 'Menu1-3' }
          }
        ]
      },
      {
        path: 'menu2',
        component: () => import('@/views/nested/menu2/index'),
        meta: { title: 'menu2' }
      }
    ]
  },

  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
        meta: { title: 'External Link', icon: 'link' }
      }
    ]
  },
  */

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
