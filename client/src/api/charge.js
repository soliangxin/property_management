import request from '@/utils/request'

export function getChangeItemList(params) {
  return request({
    url: '/change/itemList',
    method: 'get',
    params
  })
}

export function getChangeRecordList(params) {
  return request({
    url: '/change/recordList',
    method: 'get',
    params
  })
}
