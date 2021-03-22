import Vue from 'vue'

export default (context, inject) => {
  const $eventHub = new Vue()
  inject('eventHub', $eventHub)
}