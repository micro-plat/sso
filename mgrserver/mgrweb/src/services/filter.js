
import Vue from "vue"
import DateConvert from './date'

Vue.filter('DateFilter', (value, format) => {
  let res
  if (value === '') {
    return '-'
  } else {
    res = DateConvert(format, value)
    return res
  }
})

Vue.filter('StringFilter', value => {
  if (value === '') {
    return '---'
  }else{
    return value
  }
})

Vue.filter('EllipsisFilter', (value, number) => {
  if (value) {
    if (value.length <= number) {
      return value
    }
    else {
      let subval = value.slice(0, number - 1) + '...'
      return subval
    }
  }
  else {
    return '-'
  }
})
