function booleanString(b) {
	return (b ? "true" : "false");
}


function commonSuffixLength(strings) {
	if (strings.length == 0) {
		return 0;
	}
	if (strings.length == 1) {
		return 0;
	}
	var longestSuffixLength = 0;
	var maxLength = 0;
	strings.forEach(function(s) {
		maxLength = ((maxLength == 0) ? s.length : Math
				.min(maxLength, s.length));
	});
	var suffixLength = 0;
	while (suffixLength < maxLength) {
		suffixLength++
		var suffixes = strings.map(function(s) {
			return s.substring(s.length - suffixLength)
		});
		var uniqueSuffixes = suffixes.filter(function(elem, pos) {
			return suffixes.indexOf(elem) == pos;
		})
		if (uniqueSuffixes.length > 1) {
			// lost it. keep last longestSuffixLength value
			break;
		}
		// we're still good
		longestSuffixLength = suffixLength;
	}
	return longestSuffixLength;
}


function addAlert(alertText) {
	$("#alerts_container")
			.append(
					'<div class="alert alert-danger alert-dismissable">'
							+ '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>'
							+ alertText + '</div>');
	$(".alert").alert();
}
