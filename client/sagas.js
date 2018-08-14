import { takeLatest } from 'redux-saga/effects';
// import Api from '...'

function* fetchUser(action) {
    try {
        yield put({type: 'USER_FETCH_SUCCEEDED', user: {name:'quan'}});
    } catch (e) {
        yield put({type: 'USER_FETCH_FAILED', message: e.message});
    }
}

function* root() {
    yield takeLatest('USER_FETCH_REQUESTED', fetchUser);
}
export default root;