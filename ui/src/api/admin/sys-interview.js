import request from '@/utils/request'

// 查询岗位列表
export function listInterview(query) {
  return request({
    url: '/api/v1/interview',
    method: 'get',
    params: query
  })
}

// 查询岗位详细
export function getInterview(postId) {
  return request({
    url: '/api/v1/interview/' + postId,
    method: 'get'
  })
}

// 新增岗位
export function addInterview(data) {
  return request({
    url: '/api/v1/interview',
    method: 'post',
    data: data
  })
}

// 修改岗位
export function updateInterview(data, id) {
  return request({
    url: '/api/v1/interview/' + id,
    method: 'put',
    data: data
  })
}

// 删除岗位
export function delInterview(postId) {
  return request({
    url: '/api/v1/post',
    method: 'delete',
    data: postId
  })
}

