import request from '@/utils/request'

export function getStallList(params) {
  return request({
    url: '/stall/list',
    method: 'get',
    params
  })
}
