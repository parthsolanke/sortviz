(function () {
    'use strict';

    var ALGORITHMS = [
        { value: 'selection', label: 'Selection Sort' },
        { value: 'bubble', label: 'Bubble Sort' },
        { value: 'insertion', label: 'Insertion Sort' },
        { value: 'merge', label: 'Merge Sort' },
        { value: 'quick', label: 'Quick Sort' },
        { value: 'heap', label: 'Heap Sort' },
        { value: 'shell', label: 'Shell Sort' },
        { value: 'cocktail', label: 'Cocktail Shaker Sort' }
    ];

    var NAV_LINKS = [
        { key: 'visualizer', href: '/angular', label: 'Visualizer' },
        { key: 'algorithms', href: '/angular/algorithms', label: 'Algorithms' },
        { key: 'about', href: '/angular/about', label: 'About' }
    ];

    angular.module('sortvizApp', [])
        .constant('ALGORITHMS', ALGORITHMS)
        .directive('sortvizNav', function () {
            return {
                restrict: 'E',
                scope: {
                    page: '@'
                },
                template:
                    '<nav class="site-nav">' +
                        '<a class="brand" href="/">🔮 SortViz</a>' +
                        '<div class="nav-links">' +
                            '<a ng-repeat="link in links" ng-href="{{link.href}}" ' +
                                'ng-class="{ active: page === link.key }">{{link.label}}</a>' +
                        '</div>' +
                    '</nav>',
                link: function (scope) {
                    scope.links = NAV_LINKS;
                }
            };
        })
        .controller('VisualizerController', function ($scope, $http, $timeout, ALGORITHMS) {
            $scope.algorithms = ALGORITHMS;
            $scope.bars = [];
            for (var i = 1; i <= 40; i++) {
                $scope.bars.push(i);
            }
            $scope.algorithm = 'selection';
            $scope.speedValue = 6;
            $scope.steps = null;
            $scope.stepIndex = -1;
            $scope.playing = false;
            $scope.status = '';
            $scope.modal = null;

            var timer = null;

            function delay() {
                return (11 - $scope.speedValue) * 100;
            }

            function cancelTimer() {
                if (timer) {
                    $timeout.cancel(timer);
                    timer = null;
                }
            }

            function resetRun() {
                $scope.playing = false;
                cancelTimer();
                $scope.steps = null;
                $scope.stepIndex = -1;
            }

            function scheduleStep() {
                if (!$scope.playing || !$scope.steps) {
                    return;
                }
                if ($scope.stepIndex >= $scope.steps.length - 1) {
                    $scope.playing = false;
                    $scope.status = 'Array sorted!';
                    return;
                }
                timer = $timeout(function () {
                    $scope.stepIndex += 1;
                    scheduleStep();
                }, delay());
            }

            $scope.displayBars = function () {
                if ($scope.steps && $scope.stepIndex >= 0) {
                    return $scope.steps[$scope.stepIndex].array;
                }
                return $scope.bars;
            };

            $scope.barHeight = function (value) {
                return { height: (value * 9.5) + 'px' };
            };

            $scope.start = function () {
                $scope.status = '';
                $http.post('/sort', { array: $scope.bars, algorithm: $scope.algorithm })
                    .then(function (response) {
                        var data = response.data;
                        if (!angular.isArray(data) || data.length === 0) {
                            throw new Error('No sort steps returned');
                        }
                        $scope.steps = data;
                        $scope.stepIndex = 0;
                        $scope.playing = true;
                        scheduleStep();
                    })
                    .catch(function (error) {
                        var message = error && error.message ? error.message : 'Request failed';
                        $scope.status = 'Error: ' + message;
                    });
            };

            $scope.togglePlay = function () {
                if ($scope.playing) {
                    $scope.playing = false;
                    cancelTimer();
                    return;
                }
                if ($scope.steps && $scope.stepIndex >= $scope.steps.length - 1) {
                    $scope.status = 'Array already sorted! Please shuffle and try again.';
                    return;
                }
                if ($scope.steps) {
                    $scope.playing = true;
                    scheduleStep();
                    return;
                }
                $scope.start();
            };

            $scope.shuffle = function () {
                var next = $scope.bars.slice();
                for (var i = next.length - 1; i > 0; i--) {
                    var j = Math.floor(Math.random() * (i + 1));
                    var temp = next[i];
                    next[i] = next[j];
                    next[j] = temp;
                }
                $scope.bars = next;
                resetRun();
                $scope.status = '';
            };

            $scope.onAlgorithmChange = function () {
                resetRun();
                $scope.status = '';
            };

            function fetchAlgorithmData() {
                return $http.get('/static/data/algorithm_data.json').then(function (response) {
                    return response.data[$scope.algorithm];
                });
            }

            $scope.showComplexity = function () {
                fetchAlgorithmData().then(function (data) {
                    $scope.modal = { title: 'Complexity', text: data.complexity };
                }).catch(function () {
                    $scope.status = 'Error: Could not load algorithm data';
                });
            };

            $scope.showCode = function () {
                fetchAlgorithmData().then(function (data) {
                    $scope.modal = {
                        title: 'Code',
                        code: data.code.replace(/\{\{NEWLINE\}\}/g, '\n')
                    };
                }).catch(function () {
                    $scope.status = 'Error: Could not load algorithm data';
                });
            };

            $scope.closeModal = function () {
                $scope.modal = null;
            };

            $scope.$on('$destroy', cancelTimer);
        })
        .controller('AlgorithmsController', function ($scope, $http, ALGORITHMS) {
            $scope.algorithms = ALGORITHMS;
            $scope.data = null;
            $scope.error = '';

            $http.get('/static/data/algorithm_data.json')
                .then(function (response) {
                    $scope.data = response.data;
                })
                .catch(function () {
                    $scope.error = 'Could not load algorithm data';
                });

            $scope.code = function (info) {
                return info.code.replace(/\{\{NEWLINE\}\}/g, '\n');
            };
        });
})();
