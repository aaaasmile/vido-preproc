import Throttle from './throttle.js'
export default {
	debounce ( delay, atBegin, callback ) {
		return callback === undefined ? Throttle.throttle(delay, atBegin, false) : Throttle.throttle(delay, callback, atBegin !== false);
	}
}