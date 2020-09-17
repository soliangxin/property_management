import request from '@/utils/request'

export function getCommunityList(params) {
  return request({
    url: '/community/list',
    method: 'get',
    params
  })
}
