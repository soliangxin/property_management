import request from '@/utils/request'

export function getTenementList(params) {
  return request({
    url: '/tenement/list',
    method: 'get',
    params
  })
}
