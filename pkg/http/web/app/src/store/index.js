import { createStore } from 'vuex'
import axios from 'axios'

axios.defaults.baseURL = 'http://otus-social.knik.space'
axios.defaults.headers.post['Content-Type'] = 'application/json';

const token = localStorage.getItem('token')
if (token !== "") {
    axios.defaults.headers.common['Authorization'] = "Bearer " + token
}

export default createStore({
    state: {
        status: '',
        token: localStorage.getItem('token') || '',
        user: JSON.parse(localStorage.getItem('user')) || {
            id: null,
            email: '',
            name: '',
            surname: '',
            age: '',
            sex: '',
            city: '-',
            interests: [],
            friends: [],
        },
    },
    mutations: {
        auth_request(state) {
            state.status = 'loading'
        },
        auth_success(state, {token, user}) {
            state.status = 'success'
            state.token = token
            state.userID = user.id
            state.user = user
        },
        auth_error(state) {
            state.status = 'error'
        },
        logout(state) {
            state.status = ''
            state.token = ''
            state.user = {
                id: null,
                email: '',
                name: '',
                surname: '',
                age: '',
                sex: '',
                city: '--',
                interests: [],
                friends: [],
            }
        },
        profile(state, data) {
            state.user = data
        },
        append_friends(state, id) {
            if (state.user.friends == null) {
                state.user.friends = [id]
            } else {
                state.user.friends.push(id)
            }

            localStorage.setItem('user', JSON.stringify(state.user))
        }
    },
    actions: {
        login({ commit }, user) {
            return new Promise((resolve, reject) => {
                commit('auth_request')
                axios({ url: '/sign-in', data: user, method: 'POST', mode: 'no-cors'})
                    .then(resp => {
                        const token = resp.data.token
                        const user = {
                            id: resp.data.user.id,
                            email: resp.data.user.email,
                            name: resp.data.user.name,
                            surname: resp.data.user.surname,
                            age: resp.data.user.age,
                            sex: resp.data.user.sex,
                            city: resp.data.user.city,
                            interests: resp.data.user.interests,
                            friends: resp.data.user.friends,
                        }

                        localStorage.setItem('token', token)
                        localStorage.setItem('user', JSON.stringify(user))

                        axios.defaults.headers.common['Authorization'] = "Bearer " + token

                        commit('auth_success', {token, user})
                        resolve(resp)
                    })
                    .catch(err => {
                        commit('auth_error')
                        localStorage.removeItem('token')
                        reject(err)
                    })
            })
        },
        register({ commit }, user) {
            return new Promise((resolve, reject) => {
                commit('auth_request')
                axios({ url: '/sign-up', data: user, method: 'POST', mode: 'no-cors'})
                    .then(resp => {
                        const token = resp.data.token
                        const user = {
                            id: resp.data.user.id,
                            email: resp.data.user.email,
                            name: resp.data.user.name,
                            surname: resp.data.user.surname,
                            age: resp.data.user.age,
                            sex: resp.data.user.sex,
                            city: resp.data.user.city,
                            interests: resp.data.user.interests,
                            friends: resp.data.user.friends,
                        }

                        localStorage.setItem('token', token)
                        localStorage.setItem('user', JSON.stringify(user))

                        axios.defaults.headers.common['Authorization'] = "Bearer " + token

                        commit('auth_success', {token, user})
                        resolve(resp)
                    })
                    .catch(err => {
                        commit('auth_error', err)
                        localStorage.removeItem('token')
                        reject(err)
                    })
            })
        },
        logout({ commit }) {
            return new Promise((resolve) => {
                commit('logout')

                localStorage.removeItem('token')
                localStorage.removeItem('user')

                delete axios.defaults.headers.common['Authorization']

                resolve()
            })
        },
        profile({ commit }) {
            return new Promise((resolve) => {
                axios({ url: '/profile', method: 'GET', mode: 'no-cors'})
                    .then(resp => {
                        commit('profile', {
                            id: resp.data.id,
                            email: resp.data.email,
                            name: resp.data.name,
                            surname: resp.data.surname,
                            age: resp.data.age,
                            sex: resp.data.sex,
                            city: resp.data.city,
                            interests: resp.data.interests,
                            friends: resp.data.friends ?? [],
                        })
                        resolve(resp)
                    })
                    .catch(err => {
                        console.log(err)
                    })
            })
        },
        follow({ commit }, id) {
            return new Promise((resolve, reject) => {
                const data = {
                    id: id
                }
                axios({url: '/profile/friends', data: data, method: 'PUT'})
                    .then(resp => {
                        commit('append_friends', id)
                        resolve(resp)
                    })
                    .catch(err => {
                        commit('auth_error')
                        localStorage.removeItem('token')
                        reject(err)
                    })
            })
        }
    },
    getters: {
        isLoggedIn: state => !!state.token,
        authStatus: state => state.status,
        profile: state => state.user,
        userID: state => state.userID,
    }
})
