import request from '@/utils/request'

export function getWatchList(params) {
  return request({
    url: '/watch/list',
    method: 'get',
    params
  })
}
