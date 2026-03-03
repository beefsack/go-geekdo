package geekdo

var advSearchExpected = []SearchCollectionItem{
	{ID: 72125, Type: "boardgame", Rank: 101, Thumbnail: "https://cf.geekdo-images.com/cnFppsVNOSTJ-W3APQFuTg__micro/img/dfQTNAju5CUxn6JQl-Pxubac99o=/fit-in/64x64/filters:strip_icc()/pic1974056.jpg", Name: "Eclipse: New Dawn for the Galaxy", URL: "/boardgame/72125/eclipse-new-dawn-for-the-galaxy", Year: 2011, BayesAverage: 7.59, Average: 7.82, UsersRated: 28997},
	{ID: 359871, Type: "boardgame", Rank: 102, Thumbnail: "https://cf.geekdo-images.com/XWImAu_3RK61wbzcKboVdA__micro/img/LqttudgwlhjRyem9gXvYuhVaZzA=/fit-in/64x64/filters:strip_icc()/pic8145530.png", Name: "Arcs", URL: "/boardgame/359871/arcs", Year: 2024, BayesAverage: 7.585, Average: 8.02, UsersRated: 15912},
}

var advSearchPage = []byte(`<!DOCTYPE html>
<html ng-app="GeekApp" lang="en-US" ng-cloak>
<head>
	<meta charset='utf-8'>
	<meta id="vp" name="viewport" content="width=device-width, initial-scale=1.0">
			<script>
			window.addEventListener( 'DOMContentLoaded',  function() {
				var width = document.documentElement.clientWidth || window.innerWidth;
				if (width < 960) {
					var mvp = document.getElementById('vp');
					// android debugging
					mvp.setAttribute('content','width=960');
				}
			});
		</script>
		<meta content='yes' name='apple-mobile-web-app-capable'>
	<meta content='IE=edge,chrome=1' http-equiv='X-UA-Compatible'>

			<title>Browse Board Games | BoardGameGeek</title>
	
	
<link rel="apple-touch-icon" 	href="https://cf.geekdo-static.com/icons/touch-icon180.png" />
<link rel="shortcut icon" 		href="https://cf.geekdo-static.com/icons/favicon2.ico" type="image/ico" />
<link rel="icon" 					href="https://cf.geekdo-static.com/icons/favicon2.ico" type="image/ico" />
<link rel="search" 				href="/game-opensearch.xml" type="application/opensearchdescription+xml" title="BGG Game Search" />
<meta name="apple-mobile-web-app-title" content="BGG">

<meta name="theme-color" content="#2e2b47">
<link rel="preconnect" href="https://api.geekdo.com" />




	<meta property="og:image" content="https://cf.geekdo-static.com/images/opengraph/bgg_opengraph.png" />



	<meta name="keywords" content="board game, boardgames, boardgame, board, games, game, hobby, boardgamegeek, geek, geekdo">

<script>window.AdSlots = window.AdSlots || {
	cmd: [],
	disableScripts: ['gpt'],
	renderOnFirstLoad: false,
	divCheck: false
};</script>


						<link rel='stylesheet' type='text/css' href='https://cf.geekdo-static.com/static/geekui_master2_69a6444a5d900.css'>
					<link rel='stylesheet' type='text/css' href='https://cf.geekdo-static.com/static/css_master2_69a644395401a.css'>
			
	<base href="/">

<script>
	var GEEK = GEEK || {};
	GEEK.adBlock = {"blockleaderboard":0,"blocklargeleaderboard":0,"blockskyscraper":0,"blockrectangle":0,"blockitembanner":0,"blocktextad":0,"blockebay":0,"blocksupportdrive":0,"blockamazon":0,"blockamazonsearch":0,"blockcontestsmodule":0,"blockbggstore":0,"blockbuyacopy":0,"blockpreview":0,"blockbetweenposts":0,"blocksleevesad":0};
	GEEK.adConfig = {"blockleaderboard":null,"noadsense":false,"blocksupportdrive":null};
	GEEK.adSlots = {"dfp-leaderboard":{"name":"boardgame_leaderboard_728x90"},"dfp-skyscraper":{"name":"boardgame_skyscraper_160x600"},"dfp-medrect":{"name":"boardgame_rectangle_300x250"},"dfp-repeater":{"name":"boardgame_home_repeater"},"dfp-medrect-reserved-home":{"name":"boardgame_reserved_home_300x250"},"dfp-leaderboard-lg":{"name":"boardgame_home_hero"},"dfp-home-sidekick":{"name":"boardgame_home_repeater"},"dfp-inline-post":{"name":"boardgame_inline_post"},"dfp-gamepage-marketplace":{"name":"boardgame_inline_post"}};
	GEEK.legacyAds = [];
	GEEK.bggStoreAds = [];
	GEEK.googleTargets = [];
	GEEK.userid = 3281431;
	GEEK.domainname = 'boardgamegeek.com';
	GEEK.domain = 'boardgame';
	GEEK.geekitemPreload = {};
	GEEK.geekitemSettings = null;
	GEEK.geekitemModules = null;
	GEEK.geekGlobalSettings = {"shutdown_file_ops":"1","shutdown_storeimage_ops":"0","shutdown_edit_avatar":"0","shutdown_file_upload":"0","shutdown_file_download":"0","shutdown_image_upload":"0"};
	GEEK.geekimageSettings = null;
	GEEK.legacy = 1;
	GEEK.apiurlsPrefix = 'https://api.geekdo.com';
	
	
	
		
		GEEK.apiurls = {
		'root': '/api',
		'amazon': '/api/amazon',
		'files': '/api/files',
		'geekitems': '/api/geekitems',
		'collectionstatsgraph': '/api/collectionstatsgraph',
		'images': '/api/images',
		'threads': '/api/forums/threads',
	 	'forums': '/api/forums',
		'videos': '/api/videos',
		'hotness': '/api/hotness',
		'dynamicinfo': '/api/dynamicinfo',
		'subtypeinfo': '/api/subtypeinfo',
		'geekbay': '/api/geekbay',
		'geekmarket': '/market/api/v1',
	    'geekmarketapi': '/api/market',
		'geeklists': '/api/geeklists',
		'reviews': '/api/forumreviews',
		'collections': '/api/collections',
		'linkeditems': '/api/geekitem/linkeditems',
		'subscriptions': '/api/subscriptions',
	    'fans': '/api/fans',
	    'geekpreviews': '/api/geekpreviews',
	    'geekpreviewitems': '/api/geekpreviewitems',
	   	'geekpreviewparentitems': '/api/geekpreviewparentitems',
		'recs': '/api/geekitem/recs',
	    'awards': '/api/geekawards',
		'historicalrankgraph':  '/api/historicalrankgraph',
		'blueprint_recipes': '/api/blueprints/recipes',
		'affiliateads' : '/api/affiliateads',
		'sleevesbycard': '/api/sleevesbycard',
		'cardsetsbygame': '/api/cardsetsbygame',
	};
</script>

<script async src="https://www.googletagmanager.com/gtag/js?id=UA-104725-1"></script>
<script>
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('consent', 'default', {
		'ad_storage': 'denied',
		'analytics_storage': 'denied',
		'personalization_storage': 'denied',
		'wait_for_update': 3000
		 });
	gtag('config', 'UA-104725-1', {
		'cookie_domain': 'boardgamegeek.com',
		'send_page_view': false
	});
</script>

        						<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekangular17_master2_69a6445ae3165.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekui_master2_69a6444a5d900.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/js_master2_69a64438d2f9f.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekoutputtemplates_master2_69a6445bd090e.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekuicommontemplates_master2_69a6445bd8519.js'></script>
			<script>
	window.geekCookieConsent.gtagReady.then(function() {
		gtag('event', 'page_view');
	});
</script>
	<!-- why this? -->
<meta http-equiv="content-type" content="text/html;charset=UTF-8">

<!-- Add to main?
<link rel="apple-touch-icon" 	href="https://cf.geekdo-static.com/icons/appleicon.png" />
<link rel="shortcut icon" 		href="https://cf.geekdo-static.com/icons/favicon.ico" type="image/ico" />
<link rel="icon" 					href="https://cf.geekdo-static.com/icons/favicon.ico" type="image/ico" />
<link rel="search" 				href="/game-opensearch.xml" type="application/opensearchdescription+xml" title="BGG Game Search" />
-->





<!--this is used only for GeekMap -->


<!-- deal with adsense
	<script type="text/javascript">
		window.google_analytics_uacct = "UA-104725-1";
		var googletag = googletag || { };
		googletag.cmd = googletag.cmd || [];
		( function() {
			var gads = document.createElement('script');
			gads.async = true;
			gads.type  = 'text/javascript';
			gads.src   = "//www.googletagservices.com/tag/js/gpt.js";
			var node = document.getElementsByTagName('script')[0];
			node.parentNode.insertBefore(gads, node);
		} )();

		adunits = [
			{
				name:    'boardgame_button_120x45',
				size:    [ 120, 45 ],
				target:  'dfp-button'
			},
			{
				name:    'boardgame_leaderboard_728x90',
				size:    [ 728, 90 ],
				target:  'dfp-leaderboard'
			},
			{
				name:    'boardgame_skyscraper_160x600',
				size:    [ 160, 600 ],
				target:  'dfp-skyscraper'
			},
			{
				name:    'boardgame_rectangle_300x250',
				size:    [ 300, 250 ],
				target:  'dfp-medrect'
			},
			{
				name:    'boardgame_rectangle_300x250',
				size:    [ 300, 250 ],
				target:  'dfp-medrect-reserved-home'
			},
			{
				name:    'boardgame_button_120x90',
				size:    [ 120, 90 ],
				target:  'dfp-giftguide'
			}
		];

		googletag.cmd.push(function() {
			for( var i=0; i< adunits.length; i++ )
			{
				unit = adunits[i];
				googletag.defineUnit('/1005854/ca-pub-7206761047394129/'+unit.name, unit.size, unit.target).
				addService(googletag.pubads());
			}

			
			googletag.pubads().setTargeting( "subdomain", "all" );
			googletag.enableServices();
		} );

	</script>
-->

<script>
	function start() {
		
	}

	function ondomready() {
				Amazon_LoadAds();
				
		GEEK.addHandlers();
		GEEK.recaptchaKey = '6Ldyr6EaAAAAAE0F2hYgtHHqbF6nKPTENAwo6SyU';
	}

	if( typeof window.addEvent === "function" ) {
		window.addEvent('domready', function() {
			ondomready();
		} );
	} else {
		window.addEventListener( "DOMContentLoaded", function() {
			ondomready();
	} );
	}
	window.onload = start;
</script>


</head>

<body ng-controller="GeekOutput_LayoutCtrl as layoutctrl"
		class="domain-boardgame"

		ng-class="{ 'has-no-max-width' : layoutctrl.geekitemSettings.fluidlayout }">


<div class='d-flex flex-column min-vh-100'>
        	<div id="global-header-outer" class='global-header-outer'
	  ng-controller="NavCtrl as navctrl"
	  click-out="navctrl.closeMobileMenu()">

	<geeknav-menu></geeknav-menu>
</div>

	 
	
	<main class='global-body flex-grow-1'
		  id="mainbody"
		  ng-class="{ 'has-overlay-sidebar': layoutctrl.showOverlaySidebar, 'has-hidden-fixed-sidebar': layoutctrl.localStorage.hideFixedSidebar }">
		
		<a id='mainbodytarget' tabindex="-1"></a>

		<!-- Home King Ad -->
		
		<div hide-ad-block="blockleaderboard" >
			<div class="advertisement-leaderboard">
				<div class='center-block' ng-dfp-ad="dfp-leaderboard"></div>
			</div>
		</div>

		<div class='global-body-content-container container-fluid'>
							<geekoutput-sidebar deactivate-overlay-sidebar='layoutctrl.deactivateOverlaySidebar' show-overlay-sidebar='layoutctrl.showOverlaySidebar'></geekoutput-sidebar>
         
			
			<div class='global-body-content pending'  ng-class="{'ready': layoutctrl.ready}">
		
				<a id='maincontenttarget' tabindex="-1"></a>
			
				<div class="support-plea-container" hide-ad-block="blocksupportdrive">
					<geekoutput-plea></geekoutput-plea>
				</div>
				<div class="legacy">
	<div id="container" class="yui-skin-sam">
		<div id="maincontent" ng-non-bindable>
			<form method='GET' action="">
<div class='infobox'>
	<div class='fr'><a href="/browse/boardgame/page/1?sort=rank&amp;sortdir=asc" target='_self' title="first page">[1]</a>&nbsp;&nbsp;<a href="/browse/boardgame/page/1?sort=rank&amp;sortdir=asc" target='_self' title="previous page"><b>Prev &laquo;</b></a>&nbsp;&nbsp;<a href="/browse/boardgame/page/1?sort=rank&amp;sortdir=asc" target='_self' title="page 1">1</a>&nbsp;,&nbsp;<b>2</b>&nbsp;,&nbsp;<a href="/browse/boardgame/page/3?sort=rank&amp;sortdir=asc" target='_self' title="page 3">3</a>&nbsp;,&nbsp;<a href="/browse/boardgame/page/4?sort=rank&amp;sortdir=asc" target='_self' title="page 4">4</a>&nbsp;,&nbsp;<a href="/browse/boardgame/page/5?sort=rank&amp;sortdir=asc" target='_self' title="page 5">5</a>&nbsp;&nbsp;<a href="/browse/boardgame/page/3?sort=rank&amp;sortdir=asc" target='_self' title="next page"><b>Next &raquo;</b></a>&nbsp;&nbsp;<a href="/browse/boardgame/page/1745?sort=rank&amp;sortdir=asc" target='_self' title="last page">[1745]</a></div>
	<div class='fl'>
			 	 		 <p class="m-0 text-muted">We may earn a commission when you buy through our links. </p>
	 	</div>

	<div class='clear'></div>
</div>
</form>

<div id='collection_status' style='position:absolute; background:red; color:white;'></div>


<div id='collection'><div class='fl sf'><span id='collection_viewlabel'style='display:none;'>&nbsp;|&nbsp;View: <span id='collection_viewname'></span></span></div>

<div style='text-align:right; margin-bottom:5px;' class='sf'>	
		</div>



<div class="table-responsive">
<table cellpadding=0 cellspacing=1 class='collection_table' width='100%' id='collectionitems'>
<tr>	
			<th class='collection_bggrating'>
					
					<a href="/browse/boardgame?sort=rank&sortdir=desc">Board Game Rank<img alt='ascending sort' border=0 hspace=5 src='https://cf.geekdo-static.com/images/collection/arrow_up.gif'></a>
				</th>
				
			<th class='collection_thumbnail'>
			<span class="sr-only">Thumbnail image</span>
		</th>
		
			<th class='collection_title'>
						
				<a href="/browse/boardgame?sort=title">Title</a>
					</th>
	
		
					
				<th class='collection_rating'>
									Your<br>Rating
		
						
		</th>
		
			<th class='collection_bggrating'>
						<a href="/browse/boardgame?sort=bggrating">Geek Rating</a>
						
		</th>
	
			<th class='collection_bggrating'>
							<a href="/browse/boardgame?sort=avgrating">Avg Rating</a>
						
		</th>
	
			<th class='collection_bggrating'>
							<a href="/browse/boardgame?sort=numvoters">Num Voters</a>
						
		</th>
	
	
	
		
	
	
	
	
	

		
			<th class='collection_status'>
			Status
					</th>
	
				<th class='collection_plays'>
							Your<br>Plays
					</th>
	
	
		
	
		
	
		
	
																					
					<th class='collection_shop'>
			Shop
		</th>
		</tr>


	
	
			 		

<tr id='row_'>		
			<td class='collection_rank' align='center' >
			<a name="101"></a>			101			
					</td>		
						
		
			<td class='collection_thumbnail'>
			<a   href="/boardgame/72125/eclipse-new-dawn-for-the-galaxy" ><img alt="Board Game: Eclipse: New Dawn for the Galaxy"   srcset="https://cf.geekdo-images.com/cnFppsVNOSTJ-W3APQFuTg__micro/img/dfQTNAju5CUxn6JQl-Pxubac99o=/fit-in/64x64/filters:strip_icc()/pic1974056.jpg, https://cf.geekdo-images.com/cnFppsVNOSTJ-W3APQFuTg__micro@2x/img/VnsMjnBioTwpPeaemudvHNAjgiU=/fit-in/128x128/filters:strip_icc()/pic1974056.jpg 2x" src="https://cf.geekdo-images.com/cnFppsVNOSTJ-W3APQFuTg__micro/img/dfQTNAju5CUxn6JQl-Pxubac99o=/fit-in/64x64/filters:strip_icc()/pic1974056.jpg"></a>
		</td>
		
			<td id='CEcell_objectname1' 
	class="collection_objectname 
						 browse"

	>
	
	<div id='status_objectname1'></div>
	<div id='results_objectname1' style='z-index:1000;' onclick=''>
		
					<a  href="/boardgame/72125/eclipse-new-dawn-for-the-galaxy"  class='primary' >Eclipse: New Dawn for the Galaxy</a>
				
							<span class='smallerfont dull'>(2011)</span>
								
	</div>

			<p class="smallefont dull" style="margin: 2px 0 0 0;">
			Build an interstellar civilization by exploration, research, conquest, and diplomacy.
		</p>
	</td>
		
		
			<td id='CEcell_rating1' class='collection_rating editfield' align=center			
		onclick="CE_EditData( 
									{
									cellid: 		'1',
									collid:  	'', 
									fieldname: 	'rating', 
									objecttype: 'thing', 
									objectid: 	'72125'
									}	);"
		
	>
	<div id='status_rating1'></div>
	<div id='results_rating1'>
		<div style='background:white;' class='rating'><span>N/A</span></div>
	</div>
</td>	
		
			<td class='collection_bggrating' align='center'>
			7.590		</td>
	
			<td class='collection_bggrating' align='center'>
			7.82		</td>
	
			<td class='collection_bggrating' align='center'>
			28997		</td>
	
	

		

		

		

		

		

   
   
		
			<td id='CEcell_status1' class='collection_status editfield' 
			
		onclick="CE_EditData( 
									{
									cellid: 		'1',
									collid:  	'', 
									fieldname: 	'status', 
									objecttype: 'thing', 
									objectid: 	'72125'
									}	);"
		
	>
	<div id='status_status1'></div>
	<div id='results_status1'>
		





	


	
	</div>
</td>
	
		
		
		
<td id='CEcell_plays1' class='collection_plays editfield'
			
		onclick="CE_NewPlay( 
			{
				quickplayid: '2', 
				objecttype:  'thing',
				objectid:    '72125', 
				results:     $('quickplay_results2'),
				inplace: 	 true,
				cellid:		 '1'
			}
			);"
	>
	<div id='quickplay_results2'>
			</div>
</td>	
		
	
		
	
		

		

														
		
				<td class='collection_shop'>
												<div class='aad' id='aad_thing_72125_textwithprices__'></div>
									</td>
		
	</tr>
					 		

<tr id='row_'>		
			<td class='collection_rank' align='center' >
			<a name="102"></a>			102			
					</td>		
						
		
			<td class='collection_thumbnail'>
			<a   href="/boardgame/359871/arcs" ><img alt="Board Game: Arcs"   srcset="https://cf.geekdo-images.com/XWImAu_3RK61wbzcKboVdA__micro/img/LqttudgwlhjRyem9gXvYuhVaZzA=/fit-in/64x64/filters:strip_icc()/pic8145530.png, https://cf.geekdo-images.com/XWImAu_3RK61wbzcKboVdA__micro@2x/img/VrVNA_1ZwVmvfs6QDjsNNE78XzE=/fit-in/128x128/filters:strip_icc()/pic8145530.png 2x" src="https://cf.geekdo-images.com/XWImAu_3RK61wbzcKboVdA__micro/img/LqttudgwlhjRyem9gXvYuhVaZzA=/fit-in/64x64/filters:strip_icc()/pic8145530.png"></a>
		</td>
		
			<td id='CEcell_objectname3' 
	class="collection_objectname 
						 browse"

	>
	
	<div id='status_objectname3'></div>
	<div id='results_objectname3' style='z-index:1000;' onclick=''>
		
					<a  href="/boardgame/359871/arcs"  class='primary' >Arcs</a>
				
							<span class='smallerfont dull'>(2024)</span>
								
	</div>

			<p class="smallefont dull" style="margin: 2px 0 0 0;">
			Rule the galaxy with your court, fleet, and cities.
		</p>
	</td>
		
		
			<td id='CEcell_rating3' class='collection_rating editfield' align=center			
		onclick="CE_EditData( 
									{
									cellid: 		'3',
									collid:  	'', 
									fieldname: 	'rating', 
									objecttype: 'thing', 
									objectid: 	'359871'
									}	);"
		
	>
	<div id='status_rating3'></div>
	<div id='results_rating3'>
		<div style='background:white;' class='rating'><span>N/A</span></div>
	</div>
</td>	
		
			<td class='collection_bggrating' align='center'>
			7.585		</td>
	
			<td class='collection_bggrating' align='center'>
			8.02		</td>
	
			<td class='collection_bggrating' align='center'>
			15912		</td>
	
	

		

		

		

		

		

   
   
		
			<td id='CEcell_status3' class='collection_status editfield' 
			
		onclick="CE_EditData( 
									{
									cellid: 		'3',
									collid:  	'', 
									fieldname: 	'status', 
									objecttype: 'thing', 
									objectid: 	'359871'
									}	);"
		
	>
	<div id='status_status3'></div>
	<div id='results_status3'>
		





	


	
	</div>
</td>
	
		
		
		
<td id='CEcell_plays3' class='collection_plays editfield'
			
		onclick="CE_NewPlay( 
			{
				quickplayid: '4', 
				objecttype:  'thing',
				objectid:    '359871', 
				results:     $('quickplay_results4'),
				inplace: 	 true,
				cellid:		 '3'
			}
			);"
	>
	<div id='quickplay_results4'>
			</div>
</td>	
		
	
		
	
		

		

														
		
				<td class='collection_shop'>
												<div class='aad' id='aad_thing_359871_textwithprices__'></div>
									</td>
		
	</tr>
					 		

	</table>
</div>

</div>

</body>
</html>`)
