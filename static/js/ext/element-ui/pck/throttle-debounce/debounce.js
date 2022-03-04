export default {
	function ( delay, atBegin, callback ) {
		return callback === undefined ? throttle(delay, atBegin, false) : throttle(delay, callback, atBegin !== false);
	}
}