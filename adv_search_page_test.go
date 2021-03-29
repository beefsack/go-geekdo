package geekdo

var advSearchExpected = []SearchCollectionItem{
	{ID: 42, Type: "boardgame", Rank: 92, Thumbnail: "https://cf.geekdo-images.com/soAzNVWglCdVBacNjoCTJw__micro/img/7MXTV2Z2rBL40blbo6F8e1RBeRo=/fit-in/64x64/filters:strip_icc()/pic2338267.jpg", Name: "Tigris & Euphrates", URL: "/boardgame/42/tigris-euphrates", Year: 1997, BayesAverage: 7.53, Average: 7.7, UsersRated: 25368},
	{ID: 19419, Type: "boardgame", Rank: 4359, Thumbnail: "https://cf.geekdo-images.com/AbS6G6Wk8Bbd84PH2xeIzA__micro/img/fKG9MauHKDSxwE3mukJKSNrxX5w=/fit-in/64x64/filters:strip_icc()/pic1512661.jpg", Name: "Euphrates & Tigris: Contest of Kings", URL: "/boardgame/19419/euphrates-tigris-contest-kings", Year: 2005, BayesAverage: 5.735, Average: 6.08, UsersRated: 1210},
}

var advSearchPage = []byte(`

<!DOCTYPE html>
<html ng-app="GeekApp" lang="en-US" ng-cloak>
<head>
	<meta charset='utf-8'>
	<meta id="vp" name="viewport" content="width=device-width, initial-scale=1.0">
			<script>
			window.addEventListener( 'DOMContentLoaded',  function() {
				var width = document.documentElement.clientWidth || window.innerWidth;
				console.log(width);
				if (width < 960) {
					var mvp = document.getElementById('vp');
					// android debugging
					console.log(mvp);
					mvp.setAttribute('content','width=960');
				}
			});
		</script>
		<meta content='yes' name='apple-mobile-web-app-capable'>
	<meta content='IE=edge,chrome=1' http-equiv='X-UA-Compatible'>

			<title>BoardGameGeek | Gaming Unplugged Since 2000</title>
	
	
<link rel="apple-touch-icon" 	href="https://cf.geekdo-static.com/icons/touch-icon180.png" />
<link rel="shortcut icon" 		href="https://cf.geekdo-static.com/icons/favicon2.ico" type="image/ico" />
<link rel="icon" 					href="https://cf.geekdo-static.com/icons/favicon2.ico" type="image/ico" />
<link rel="search" 				href="/game-opensearch.xml" type="application/opensearchdescription+xml" title="BGG Game Search" />
<meta name="apple-mobile-web-app-title" content="BGG">


<link rel="preconnect" href="https://www.google.com" />
<link rel="preconnect" href="https://api.geekdo.com" />


	<meta property="og:image" content="https://cf.geekdo-static.com/images/opengraph/bgg_opengraph.png" />




	<meta name="keywords" content="board game, boardgames, boardgame, board, games, game, hobby, boardgamegeek, geek, geekdo">


						<link rel='stylesheet' type='text/css' href='https://cf.geekdo-static.com/static/geekui_master2_605e44a69a734.css'>
					<link rel='stylesheet' type='text/css' href='https://cf.geekdo-static.com/static/css_master2_605e449a6e466.css'>
			
	<base href="/">

	<script src="https://kit.fontawesome.com/42d3f5a072.js" crossorigin="anonymous"></script>

<script>
	var GEEK = GEEK || {};
	GEEK.adBlock = [];
	GEEK.adConfig = {"blockleaderboard":true,"blockskyscraper":null,"noadsense":null};
	GEEK.adSlots = {"dfp-leaderboard":{"name":"boardgame_leaderboard_728x90"},"dfp-skyscraper":{"name":"boardgame_skyscraper_160x600"},"dfp-medrect":{"name":"boardgame_rectangle_300x250"},"dfp-repeater":{"name":"boardgame_home_repeater"},"dfp-medrect-reserved-home":{"name":"boardgame_reserved_home_300x250"},"dfp-leaderboard-lg":{"name":"boardgame_home_hero"},"dfp-home-sidekick":{"name":"boardgame_home_sidekick"}};
	GEEK.googleTargets = null;
	GEEK.userid = 0;
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
		'amazon': '/api/amazon',
		'files': '/api/files',
		'geekitems': '/api/geekitems',
		'images': '/api/images',
		'threads': '/api/forums/threads',
	 	'forums': '/api/forums',
		'videos': '/api/videos',
		'hotness': '/api/hotness',
		'dynamicinfo': '/api/dynamicinfo',
		'subtypeinfo': '/api/subtypeinfo',
		'geekbay': '/api/geekbay',
		'geekmarket': '/geekmarket/api/v1',
	    'geekmarketapi': '/api/geekmarket',
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
		'historicalrankgraph':  '/api/historicalrankgraph'
	};
</script>

							<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekangular17_master2_605e44b09ec06.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekui_master2_605e44a69a734.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/js_master2_605e449a06bf2.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekoutputtemplates_master2_605e44b152e36.js'></script>
					<script type='text/javascript' src='https://cf.geekdo-static.com/static/geekuicommontemplates_master2_605e44b15bb5f.js'></script>
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
		GEEK.recaptchaKey = '6LfpC3MUAAAAAGRUk4nA-p8bk2F_8ijWX5PIdzun';
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
		ng-class="{ 'has-no-max-width' : layoutctrl.geekitemSettings.fluidlayout }">


<div class='d-flex flex-column min-vh-100'>
	<div hide-ad-block="blockleaderboard">
		<div class="advertisement advertisement-leaderboard">
			<div class='center-block' ng-dfp-ad="dfp-leaderboard"></div>
		</div>
	</div>

	
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
		
		<div class='global-body-content-container container-fluid'>
			<!-- <button class="btn btn-primary hidden-xs btn-lg feedback-trigger"
						login-required
					  feedback-widget
					  ng-click="feedbackctrl.openFeedback()">
				<i class="fi-comments"></i>&nbsp; Feedback
			</button> -->

			<geekoutput-sidebar deactivate-overlay-sidebar='layoutctrl.deactivateOverlaySidebar' show-overlay-sidebar='layoutctrl.showOverlaySidebar'></geekoutput-sidebar>

			
			<div class='global-body-content pending'  ng-class="{'ready': layoutctrl.ready}">
				<a id='maincontenttarget' tabindex="-1"></a>
			
				<div class="legacy">
	<div id="container" class="yui-skin-sam">
		<div id="maincontent" ng-non-bindable>
			<div style="margin: auto;"><iframe src="/amazon/iframesearchad?searchterm=euphrates&position=leaderboard&subype=boardgame" class="amazon_search_ad"></iframe></div><form method='GET' action="">
<div class='infobox'>
	<div class='fr'></div>
	
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
					
					<a href="/search/boardgame?sort=rank&advsearch=1&q=euphrates&include%5Bdesignerid%5D=&include%5Bpublisherid%5D=&geekitemname=&range%5Byearpublished%5D%5Bmin%5D=&range%5Byearpublished%5D%5Bmax%5D=&range%5Bminage%5D%5Bmax%5D=&range%5Bnumvoters%5D%5Bmin%5D=&range%5Bnumweights%5D%5Bmin%5D=&range%5Bminplayers%5D%5Bmax%5D=&range%5Bmaxplayers%5D%5Bmin%5D=&range%5Bleastplaytime%5D%5Bmin%5D=&range%5Bplaytime%5D%5Bmax%5D=&floatrange%5Bavgrating%5D%5Bmin%5D=&floatrange%5Bavgrating%5D%5Bmax%5D=&floatrange%5Bavgweight%5D%5Bmin%5D=&floatrange%5Bavgweight%5D%5Bmax%5D=&colfiltertype=&searchuser=beefsack&playerrangetype=normal&B1=Submit">Board Game Rank</a>
				</th>
				
			<th class='collection_thumbnail'>
			<span class="sr-only">Thumbnail image</span>
		</th>
		
			<th class='collection_title'>
						
				<a href="/search/boardgame?sort=title&advsearch=1&q=euphrates&include%5Bdesignerid%5D=&include%5Bpublisherid%5D=&geekitemname=&range%5Byearpublished%5D%5Bmin%5D=&range%5Byearpublished%5D%5Bmax%5D=&range%5Bminage%5D%5Bmax%5D=&range%5Bnumvoters%5D%5Bmin%5D=&range%5Bnumweights%5D%5Bmin%5D=&range%5Bminplayers%5D%5Bmax%5D=&range%5Bmaxplayers%5D%5Bmin%5D=&range%5Bleastplaytime%5D%5Bmin%5D=&range%5Bplaytime%5D%5Bmax%5D=&floatrange%5Bavgrating%5D%5Bmin%5D=&floatrange%5Bavgrating%5D%5Bmax%5D=&floatrange%5Bavgweight%5D%5Bmin%5D=&floatrange%5Bavgweight%5D%5Bmax%5D=&colfiltertype=&searchuser=beefsack&playerrangetype=normal&B1=Submit">Title</a>
					</th>
	
		
		
			<th class='collection_bggrating'>
						<a href="/search/boardgame?sort=bggrating&advsearch=1&q=euphrates&include%5Bdesignerid%5D=&include%5Bpublisherid%5D=&geekitemname=&range%5Byearpublished%5D%5Bmin%5D=&range%5Byearpublished%5D%5Bmax%5D=&range%5Bminage%5D%5Bmax%5D=&range%5Bnumvoters%5D%5Bmin%5D=&range%5Bnumweights%5D%5Bmin%5D=&range%5Bminplayers%5D%5Bmax%5D=&range%5Bmaxplayers%5D%5Bmin%5D=&range%5Bleastplaytime%5D%5Bmin%5D=&range%5Bplaytime%5D%5Bmax%5D=&floatrange%5Bavgrating%5D%5Bmin%5D=&floatrange%5Bavgrating%5D%5Bmax%5D=&floatrange%5Bavgweight%5D%5Bmin%5D=&floatrange%5Bavgweight%5D%5Bmax%5D=&colfiltertype=&searchuser=beefsack&playerrangetype=normal&B1=Submit">Geek Rating</a>
						
		</th>
	
			<th class='collection_bggrating'>
							<a href="/search/boardgame?sort=avgrating&advsearch=1&q=euphrates&include%5Bdesignerid%5D=&include%5Bpublisherid%5D=&geekitemname=&range%5Byearpublished%5D%5Bmin%5D=&range%5Byearpublished%5D%5Bmax%5D=&range%5Bminage%5D%5Bmax%5D=&range%5Bnumvoters%5D%5Bmin%5D=&range%5Bnumweights%5D%5Bmin%5D=&range%5Bminplayers%5D%5Bmax%5D=&range%5Bmaxplayers%5D%5Bmin%5D=&range%5Bleastplaytime%5D%5Bmin%5D=&range%5Bplaytime%5D%5Bmax%5D=&floatrange%5Bavgrating%5D%5Bmin%5D=&floatrange%5Bavgrating%5D%5Bmax%5D=&floatrange%5Bavgweight%5D%5Bmin%5D=&floatrange%5Bavgweight%5D%5Bmax%5D=&colfiltertype=&searchuser=beefsack&playerrangetype=normal&B1=Submit">Avg Rating</a>
						
		</th>
	
			<th class='collection_bggrating'>
							<a href="/search/boardgame?sort=numvoters&advsearch=1&q=euphrates&include%5Bdesignerid%5D=&include%5Bpublisherid%5D=&geekitemname=&range%5Byearpublished%5D%5Bmin%5D=&range%5Byearpublished%5D%5Bmax%5D=&range%5Bminage%5D%5Bmax%5D=&range%5Bnumvoters%5D%5Bmin%5D=&range%5Bnumweights%5D%5Bmin%5D=&range%5Bminplayers%5D%5Bmax%5D=&range%5Bmaxplayers%5D%5Bmin%5D=&range%5Bleastplaytime%5D%5Bmin%5D=&range%5Bplaytime%5D%5Bmax%5D=&floatrange%5Bavgrating%5D%5Bmin%5D=&floatrange%5Bavgrating%5D%5Bmax%5D=&floatrange%5Bavgweight%5D%5Bmin%5D=&floatrange%5Bavgweight%5D%5Bmax%5D=&colfiltertype=&searchuser=beefsack&playerrangetype=normal&B1=Submit">Num Voters</a>
						
		</th>
	
	
	
		
	
	
	
	
	

		
	
	
	
		
	
		
	
		
	
																					
					<th class='collection_shop'>
			Shop
		</th>
		</tr>


	
	
			

<tr id='row_'>		
			<td class='collection_rank' align='center' >
			<a name="92"></a>			92			
					</td>		
						
		
			<td class='collection_thumbnail'>
			<a   href="/boardgame/42/tigris-euphrates" ><img alt="Board Game: Tigris & Euphrates"   src="https://cf.geekdo-images.com/soAzNVWglCdVBacNjoCTJw__micro/img/7MXTV2Z2rBL40blbo6F8e1RBeRo=/fit-in/64x64/filters:strip_icc()/pic2338267.jpg"></a>
		</td>
		
			<td id='CEcell_objectname1' 
	class="collection_objectname 
						 browse"

	>
	
	<div id='status_objectname1'></div>
	<div id='results_objectname1' style='z-index:1000;' onclick=''>
		
					<a  href="/boardgame/42/tigris-euphrates"  class='primary' >Tigris & Euphrates</a>
				
							<span class='smallerfont dull'>(1997)</span>
								
	</div>

			<p class="smallefont dull" style="margin: 2px 0 0 0;">
			Keep your Mesopotamian civilisation in perfect balance through revolutions and wars.
		</p>
	</td>
		
		
		
			<td class='collection_bggrating' align='center'>
			7.530		</td>
	
			<td class='collection_bggrating' align='center'>
			7.70		</td>
	
			<td class='collection_bggrating' align='center'>
			25368		</td>
	
	

		

		

		

		

		

   
   
		
		
		
	
		
	
		

		

														
		
				<td class='collection_shop'>
												<div class='aad' id='aad_thing_42_textwithprices__'></div>
																	[<a href="/boardgame/42/tigris-euphrates/marketplace/geekmarket">Shop</a>]
					</td>
		
	</tr>
					

<tr id='row_'>		
			<td class='collection_rank' align='center' >
			<a name="4359"></a>			4359			
					</td>		
						
		
			<td class='collection_thumbnail'>
			<a   href="/boardgame/19419/euphrates-tigris-contest-kings" ><img alt="Board Game: Euphrates & Tigris: Contest of Kings"   src="https://cf.geekdo-images.com/AbS6G6Wk8Bbd84PH2xeIzA__micro/img/fKG9MauHKDSxwE3mukJKSNrxX5w=/fit-in/64x64/filters:strip_icc()/pic1512661.jpg"></a>
		</td>
		
			<td id='CEcell_objectname2' 
	class="collection_objectname 
						 browse"

	>
	
	<div id='status_objectname2'></div>
	<div id='results_objectname2' style='z-index:1000;' onclick=''>
		
					<a  href="/boardgame/19419/euphrates-tigris-contest-kings"  class='primary' >Euphrates & Tigris: Contest of Kings</a>
				
							<span class='smallerfont dull'>(2005)</span>
								
	</div>

	</td>
		
		
		
			<td class='collection_bggrating' align='center'>
			5.735		</td>
	
			<td class='collection_bggrating' align='center'>
			6.08		</td>
	
			<td class='collection_bggrating' align='center'>
			1210		</td>
	
	

		

		

		

		

		

   
   
		
		
		
	
		
	
		

		

														
		
				<td class='collection_shop'>
												<div class='aad' id='aad_thing_19419_textwithprices__'></div>
																	[<a href="/boardgame/19419/euphrates-tigris-contest-kings/marketplace/geekmarket">Shop</a>]
					</td>
		
	</tr>
		

	</table>
</div>

</div>

<p align='right'></p>

			<div id="legacy_modal"></div>
		</div>
	</div>
</div>

				<div class='global-body-content-secondary'>
									</div>
			</div>

			<div class='global-body-ad' hide-ad-block="blockskyscraper">
				<div class='advertisement advertisement-top'>
					<div class="center-block advertisement-adsense-skyscraper" ng-dfp-ad="dfp-skyscraper"></div>
										<iframe src="/amazon/iframeskyscraperad" class="amazon_skyscraper_ad" title="Amazon ad"></iframe>
				</div>
			</div>

		</div>

	</main>
			<geekoutput-footer></geekoutput-footer>
	</div>


<script>
	window.google_analytics_uacct = "UA-104725-1";
	(function ( i, s, o, g, r, a, m ) {
		i[ 'GoogleAnalyticsObject' ] = r;
		i[ r ] = i[ r ] || function () {
				(i[ r ].q = i[ r ].q || []).push( arguments )
			}, i[ r ].l = 1 * new Date();
		a = s.createElement( o ), m = s.getElementsByTagName( o )[ 0 ];
		a.async = 1;
		a.src = g;
		m.parentNode.insertBefore( a, m )
	})
	( window, document, 'script', '//www.google-analytics.com/analytics.js', 'ga' );
	ga( 'create', 'UA-104725-1', 'auto' );
	ga( 'require', 'displayfeatures' );

	</script>

<script>
	ga( 'send', 'pageview' );
</script>



</body>
</html>`)
