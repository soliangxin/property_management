import request from '@/utils/request'

export function getHouseList(params) {
  return request({
    url: '/house/list',
    method: 'get',
    params
  })
}
