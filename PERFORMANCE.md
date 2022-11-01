# Performance optimization history

## Initial (very naive) implementation, halfway step 6

### Bench

```
goos: linux
goarch: amd64
pkg: github.com/vgalaktionov/roguelike-go/game
cpu: AMD EPYC Processor
BenchmarkRandomGameWithRoomsAndCorridors-3             1        20705952463 ns/op       179078800 B/op    738364 allocs/op
PASS
ok      github.com/vgalaktionov/roguelike-go/game       20.855s
```

### Top CPU functions:

```
Showing nodes accounting for 5840ms, 69.28% of 8430ms total
Dropped 92 nodes (cum <= 42.15ms)
Showing top 10 nodes out of 51
      flat  flat%   sum%        cum   cum%
     960ms 11.39% 11.39%     2420ms 28.71%  github.com/vgalaktionov/roguelike-go/draw.ConsoleRenderer.SetCellContent
     890ms 10.56% 21.95%      890ms 10.56%  github.com/vgalaktionov/roguelike-go/draw.ColorFromPalette
     700ms  8.30% 30.25%      700ms  8.30%  github.com/gdamore/tcell/v2.(*CellBuffer).GetContent
     630ms  7.47% 37.72%     1010ms 11.98%  runtime.mapaccess2
     550ms  6.52% 44.25%      560ms  6.64%  github.com/gdamore/tcell/v2.(*CellBuffer).SetContent
     480ms  5.69% 49.94%     1460ms 17.32%  github.com/gdamore/tcell/v2.(*simscreen).drawCell
     470ms  5.58% 55.52%     1110ms 13.17%  github.com/gdamore/tcell/v2.(*simscreen).SetContent
     450ms  5.34% 60.85%     3660ms 43.42%  github.com/vgalaktionov/roguelike-go/game/systems.clearMap
     360ms  4.27% 65.12%     5070ms 60.14%  github.com/vgalaktionov/roguelike-go/game/systems.RenderMap
     350ms  4.15% 69.28%      350ms  4.15%  github.com/vgalaktionov/roguelike-go/game/resources.(*Map).ClearContents
```

## Top memory functions

```
Showing nodes accounting for 165.21MB, 99.02% of 166.85MB total
Dropped 21 nodes (cum <= 0.83MB)
Showing top 10 nodes out of 29
      flat  flat%   sum%        cum   cum%
  129.94MB 77.88% 77.88%   129.94MB 77.88%  github.com/norendren/go-fov/fov.(*View).fov
   13.05MB  7.82% 85.70%    13.05MB  7.82%  github.com/gdamore/tcell/v2.(*CellBuffer).Resize
    7.55MB  4.53% 90.23%    20.60MB 12.35%  github.com/gdamore/tcell/v2.(*simscreen).SetSize
       5MB  3.00% 93.22%        5MB  3.00%  github.com/gdamore/tcell/v2.(*simscreen).drawCell
       3MB  1.80% 95.02%        3MB  1.80%  github.com/vgalaktionov/roguelike-go/ecs.QueryEntitiesIter
    2.51MB  1.51% 96.53%     2.51MB  1.51%  github.com/vgalaktionov/roguelike-go/game/resources.NewEmptyMap
       2MB  1.20% 97.73%   131.94MB 79.08%  github.com/norendren/go-fov/fov.(*View).Compute
    1.16MB  0.69% 98.42%     1.16MB  0.69%  runtime/pprof.StartCPUProfile
       1MB   0.6% 99.02%        1MB   0.6%  runtime.malg
         0     0% 99.02%        5MB  3.00%  github.com/gdamore/tcell/v2.(*simscreen).Show
```

## After optimizing color lookups:

### Bench

```
goos: linux
goarch: amd64
pkg: github.com/vgalaktionov/roguelike-go/game
cpu: AMD EPYC Processor
BenchmarkRandomGameWithRoomsAndCorridors-3             1        19016077564 ns/op       178930904 B/op    717044 allocs/op
PASS
ok      github.com/vgalaktionov/roguelike-go/game       19.234s
```

### Top CPU functions:

```
Showing nodes accounting for 5250ms, 78.48% of 6690ms total
Dropped 107 nodes (cum <= 33.45ms)
Showing top 10 nodes out of 51
      flat  flat%   sum%        cum   cum%
    1140ms 17.04% 17.04%     2320ms 34.68%  github.com/vgalaktionov/roguelike-go/draw.ConsoleRenderer.SetCellContent
     660ms  9.87% 26.91%      690ms 10.31%  github.com/gdamore/tcell/v2.(*CellBuffer).SetContent
     630ms  9.42% 36.32%     1130ms 16.89%  runtime.mapaccess2
     600ms  8.97% 45.29%     1350ms 20.18%  github.com/gdamore/tcell/v2.(*simscreen).drawCell
     570ms  8.52% 53.81%      570ms  8.52%  github.com/gdamore/tcell/v2.(*CellBuffer).GetContent
     540ms  8.07% 61.88%      540ms  8.07%  aeshashbody
     400ms  5.98% 67.86%     1180ms 17.64%  github.com/gdamore/tcell/v2.(*simscreen).SetContent
     250ms  3.74% 71.60%      250ms  3.74%  github.com/vgalaktionov/roguelike-go/game/resources.(*Map).PopulateBlocked
     230ms  3.44% 75.04%      230ms  3.44%  github.com/vgalaktionov/roguelike-go/game/resources.(*Map).ClearContents
     230ms  3.44% 78.48%     3920ms 58.59%  github.com/vgalaktionov/roguelike-go/game/systems.RenderMap
```

## Top memory functions

```
Showing nodes accounting for 173.75MB, 96.92% of 179.28MB total
Dropped 26 nodes (cum <= 0.90MB)
Showing top 10 nodes out of 33
      flat  flat%   sum%        cum   cum%
  138.47MB 77.24% 77.24%   138.47MB 77.24%  github.com/norendren/go-fov/fov.(*View).fov
   13.05MB  7.28% 84.51%    13.05MB  7.28%  github.com/gdamore/tcell/v2.(*CellBuffer).Resize
    7.55MB  4.21% 88.73%    20.60MB 11.49%  github.com/gdamore/tcell/v2.(*simscreen).SetSize
    3.52MB  1.97% 90.69%     3.52MB  1.97%  github.com/vgalaktionov/roguelike-go/game/resources.NewEmptyMap
    3.50MB  1.95% 92.65%   141.97MB 79.19%  github.com/norendren/go-fov/fov.(*View).Compute
    2.50MB  1.39% 94.04%     2.50MB  1.39%  github.com/gdamore/tcell/v2.(*simscreen).drawCell
    1.50MB  0.84% 94.88%     1.50MB  0.84%  github.com/vgalaktionov/roguelike-go/ecs.QueryEntitiesIter
    1.50MB  0.84% 95.72%        2MB  1.12%  github.com/vgalaktionov/roguelike-go/game/systems.MapIndexing
    1.16MB  0.65% 96.36%     1.16MB  0.65%  runtime/pprof.StartCPUProfile
       1MB  0.56% 96.92%     1.50MB  0.84%  github.com/vgalaktionov/roguelike-go/game/systems.RenderMap
```
