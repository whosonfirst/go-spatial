var whosonfirst = whosonfirst || {};
whosonfirst.spatial = whosonfirst.spatial || {};

whosonfirst.spatial.api = (function(){

    var self = {

	'point_in_polygon': function(args, on_success, on_error) {

	    var lat = args['latitude'];
	    var lon = args['longitude'];

	    var rel_url = "/intersects?latitude=" + lat + "&longitude=" + lon;
	    return self.get(rel_url, on_success, on_error);
	},

	'point_in_polygon_candidates': function(args, on_success, on_error) {

	    var lat = args['latitude'];
	    var lon = args['longitude'];

	    var rel_url = "/intersects/candidates?latitude=" + lat + "&longitude=" + lon;
	    return self.get(rel_url, on_success, on_error);
	},
	
	'get': function(rel_url, on_success, on_error) {
	    
	    var req = new XMLHttpRequest();
	    
	    req.onload = function(){
		
		var rsp;
		
		try {
		    rsp = JSON.parse(this.responseText);
            	}
		
		catch (e){
		    console.log("ERR", url, e);
		    on_error(e);
		    return false;
		}
		
		on_success(rsp);
       	    };

	    var abs_url = self.abs_url(rel_url);
	    
	    req.open("get", abs_url, true);
	    req.send();	    
	},

	'abs_url': function(rel_url) {

	    return location.protocol + "//" + location.host + rel_url;
	}
    };

    return self;
    
})();
