export const resizeIframe = () => {
    + function($) {
        $(function() {
            const oIframe = document.getElementById('show-iframe');
            //const deviceWidth = document.documentElement.clientWidth;

            // oIframe.style.width = deviceWidth + 'px';
            oIframe.style.height = (document.body.clientHeight - 170) + 'px';

            // window.onresize = () => {
            //     const oIframe = document.getElementById('show-iframe');
            //     if (oIframe) {
            //         const deviceWidth = document.documentElement.clientWidth;
            //         const deviceHeight = document.documentElement.clientHeight;
            //         oIframe.style.width = deviceWidth + 'px';
            //         oIframe.style.height = deviceHeight + 'px';
            //     }
            //
            // };


          window.onresize = function(){
            var oIframe = document.getElementById('show-iframe');
                    // oIframe.style.width = document.documentElement.clientWidth + 'px';
                    console.log("iframe: ",oIframe.contentDocument)
                    oIframe.style.height = (document.body.clientHeight - 160) + 'px';
          }

        })
    }(jQuery)
}


export const loadNavbar = () => {
    + function($) {
        $(function() {

            // class
            $(document).on('click', '[data-toggle^="class"]', function(e) {
                e && e.preventDefault();
                var $this = $(e.target),
                    $class, $target, $tmp, $classes, $targets;
                !$this.data('toggle') && ($this = $this.closest('[data-toggle^="class"]'));
                $class = $this.data()['toggle'];
                $target = $this.data('target') || $this.attr('href');
                $class && ($tmp = $class.split(':')[1]) && ($classes = $tmp.split(','));
                $target && ($targets = $target.split(','));
                $classes && $classes.length && $.each($targets, function(index, value) {
                    if ($classes[index].indexOf('*') !== -1) {
                        var patt = new RegExp('\\s' +
                            $classes[index].replace(/\*/g, '[A-Za-z0-9-_]+').split(' ').join('\\s|\\s') +
                            '\\s', 'g');
                        $($this).each(function(i, it) {
                            var cn = ' ' + it.className + ' ';
                            while (patt.test(cn)) {
                                cn = cn.replace(patt, ' ');
                            }
                            it.className = $.trim(cn);
                        });
                    }
                    ($targets[index] != '#') && $($targets[index]).toggleClass($classes[index]) || $this.toggleClass($classes[index]);
                });
                $this.toggleClass('active');
            });

            // collapse nav
            $(document).on('click', 'nav a', function(e) {
                var $this = $(e.target),
                    $active;
                $this.is('a') || ($this = $this.closest('a'));

                $active = $this.parent().siblings(".active");
                $active && $active.toggleClass('active').find('> ul:visible').slideUp(200);

                ($this.parent().hasClass('active') && $this.next().slideUp(200)) || $this.next().slideDown(200);
                $this.parent().toggleClass('active');

                $this.next().is('ul') && e.preventDefault();

                setTimeout(function() {
                    $(document).trigger('updateNav');
                }, 300);
            });
        });
    }(jQuery)
}
