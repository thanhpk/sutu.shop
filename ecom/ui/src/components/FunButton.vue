<template>
<div style="display:inline-block">
	<svg xmlns="http://www.w3.org/2000/svg" version="1.1" class="svg-filters">
						<defs>
							<filter class="filter-goo-2">
								<feGaussianBlur in="SourceGraphic" stdDeviation="7" result="blur" x="0"></feGaussianBlur>
								<feColorMatrix in="blur" mode="matrix" values="1 0 0 0 0  0 1 0 0 0  0 0 1 0 0  0 0 0 19 -9" result="goo"></feColorMatrix>
								<feComposite in="SourceGraphic" in2="goo"></feComposite>
							</filter>
						</defs>
					</svg>

	<div class="button button--2">
	<slot></slot>
	<span class="button__bg" style="transform: matrix(1, 0, 0, 1, 0, 0);"></div>
</div>
</template>
<script>
	import gsap from 'gsap';
import {TimelineLite} from 'gsap';

function getRandom(min, max){
  return Math.random() * (max - min) + min;
}

	
	export default {

	mounted: function() {
	initBt2(this.$el);


	}
	
	}


		function initBt2(document) {
			var bt = document.querySelectorAll('.button--2')[0];
  var filter = document.querySelectorAll('.filter-goo-2 feGaussianBlur')[0];
  var particleCount = 12;
  var colors = ['#DE8AA0', '#8AAEDE', '#FFB300', '#60C7DA']

			bt.addEventListener('click', function() {
					
    var particles = [];
    var tl = new TimelineLite({onUpdate: function() {
      filter.setAttribute('x', 0);
    }});
    
    tl.to(bt.querySelectorAll('.button__bg'), 0.6, { scaleX: 1.05 });
    tl.to(bt.querySelectorAll('.button__bg'), 0.9, { scale: 1, ease: Elastic.easeOut.config(1.2, 0.4) }, 0.6);

    for (var i = 0; i < particleCount; i++) {
      particles.push(window['document'].createElement('span'));
      bt.appendChild(particles[i]);

      particles[i].classList.add(i % 2 ? 'left' : 'right');
      
      var dir = i % 2 ? '-' : '+';
      var r = i % 2 ? getRandom(-1, 1)*i/2 : getRandom(-1, 1)*i;
      var size = i < 2 ? 1 : getRandom(0.4, 0.8);
      var tl = new TimelineLite({ onComplete: function(i) {
        particles[i].parentNode.removeChild(particles[i]);
        this.kill();
      }, onCompleteParams: [i] });

      tl.set(particles[i], { scale: size });
      tl.to(particles[i], 0.6, { x: dir + 20, scaleX: 3, ease: SlowMo.ease.config(0.1, 0.7, false) });
      tl.to(particles[i], 0.1, { scale: size, x: dir +'=25' }, '-=0.1');
      if(i >= 2) tl.set(particles[i], { backgroundColor: colors[Math.round(getRandom(0, 3))] });
      tl.to(particles[i], 0.6, { x: dir + getRandom(60, 100), y: r*10, scale: 0.1, ease: Power3.easeOut });
      tl.to(particles[i], 0.2, { opacity: 0, ease: Power3.easeOut }, '-=0.2');
    }
  });
}



</script>


<style>

  .button--1:focus, .button--1 .button__bg:focus, .button--2:focus, .button--4:focus, .button--4 .button__bg:focus, .button--5:focus, .button--6:focus, .button--7:focus, .button--8:focus, .button--9:focus, .button--10:focus {
    outline: none;
}


	.button {
	cursor: pointer;
	display: inline-block;
	position: relative;
	filter: url('#filter-goo-2');
	text-decoration: none;
  user-select: none;

	
     }


	
.button--2 {
  outline: 90px solid transparent !important;
  position: relative;
  z-index: 0;
  background-color: transparent; }
  .button--2 .left, .button--2 .right {
    position: absolute;
    width: 25px;
    height: 25px;
    border-radius: 15px;
    background: #222;
    -webkit-transition: background 0.1s ease-out;
    -moz-transition: background 0.1s ease-out;
    transition: background 0.1s ease-out;
    top: 50%;
    margin-top: -12px;
    z-index: -2; }
    .button--2 .left.left, .button--2 .right.left {
      left: 0; }
    .button--2 .left.right, .button--2 .right.right {
  right: 0;

	}
  .button--2 .button__bg {

	border-radius: 10px;
    content: "";
    background: #444;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: -1;
    -webkit-transition: background 0.1s ease-out;
    -moz-transition: background 0.1s ease-out;
    transition: background 0.1s ease-out; }
  .button--2:hover {
    background-color: transparent; }
    .button--2:hover:before, .button--2:hover span {
      background-color: #ffd32c; }




</style>
