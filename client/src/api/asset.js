import request from '@/utils/request'

export function getAssetList(params) {
  return request({
    url: '/asset/list',
    method: 'get',
    params
  })
}
