import { applyMiddleware, createStore as reduxCreateStore } from 'redux';
import { createLogger } from 'redux-logger';
import reducers from './reducers';
import createSagaMiddleware from 'redux-saga';
const sagaMiddleware = createSagaMiddleware();
const middlewares = [sagaMiddleware];
import root from './sagas';

// Add state logger
if (process.env.NODE_ENV !== 'production') {
  try {
    middlewares.push(createLogger());
  } catch (e) {}
}

export function createStore(state) {
  return reduxCreateStore(
    reducers,
    state,
    applyMiddleware.apply(null, middlewares)
  );
}

export let store = null;
export function getStore() { return store; }
export function setAsCurrentStore(s) {
  store = s;
  store.runSagas = () => sagaMiddleware.run(root);

  if (process.env.NODE_ENV !== 'production'
    && typeof window !== 'undefined') {
    window.store = store;
  }

  return store;
}
