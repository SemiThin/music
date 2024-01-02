// 请求统一baseUrl
const baseUrl = 'https://music.demo.net'
// 封装wx.request
function request (url, data) {
  return new Promise((resolve, reject) => {
    wx.request({
      url,
      method: 'POST',
      data,
      success (res) {
        console.info('请求成功', url, res)
        resolve(res.data)
      },
      fail: e => {
        console.error('请求失败', url, e)
        wx.showToast({
          title: e.errMsg,
          icon: 'error',
          duration: 2000
        })
        reject(e)
      }
    })
  })
}
export default {
  newsong: () => request(baseUrl + '/song/all')
}
