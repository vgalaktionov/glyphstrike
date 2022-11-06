# Performance optimization history

It's good to have the performance in mind from the get go, and even this tiny game is already showing signs
of slowing down when trying to move very fast.
Let's get a simple benchmark going and improve it as we go.

We inject a special system to randomly generate player input, and extend the game engine with an alternative
run method that will exit after N steps.

## Initial (very naive) implementation, halfway RLTK step 6 in terms of features

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

We can see that a ridiculous amount of CPU is spent on instantiating color objects at runtime.
Luckily, we have a limited palette, and it's easy to precompute these in an `init()` function.
Let's see if that helps. Along the way, we can simplify some unnecessary abstractions.

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

### Top memory functions

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

Now we can see that hashmap access is taking up a large part of the turn processing.
This happens in the ECS, and will only increase exponentially with number of entities and components.
Time to fix that. Along the way, let's add tests and benchmarks for the ECS separately.

Currently we are storing entities in hashmaps, and their components in nested hashmaps.
While hashes are pretty fast, this adds up - and we can elide this overhead if we store them in arrays and access by index. In theory that should have massive cache locality benefits too.

Furthermore, we also access the component maps for the purpose of querying - but this can be avoided by using bitmap indices to track which entity has which components.

The one downside is requiring a tiny bit more boilerplate, as we need global static sequential unique integers for component ids this way. I thought about code generation with `// go:generate` but it's not worth it at this stage.
This is Go after all, typing things is part of the fun.

## Optimizing the ECS

First, let's benchmark the current most common operations. We'll fill it out as we go along.
All benchmarks use a world seeded with 1_000_000 initial entities.

```
BenchmarkAddEntity
BenchmarkAddEntity-3                     4417538               271.4 ns/op
BenchmarkHasEntity
BenchmarkHasEntity-3                    16264411                70.18 ns/op
BenchmarkAddEntityComponent
BenchmarkAddEntityComponent-3            1000000              1279 ns/op
BenchmarkHasEntityComponent
BenchmarkHasEntityComponent-3           10290847               106.0 ns/op
BenchmarkRemoveEntity
BenchmarkRemoveEntity-3                  1965645               647.3 ns/op
BenchmarkQueryEntitiesIter
BenchmarkQueryEntitiesIter-3                   3         424827922 ns/op
BenchmarkQueryEntitiesSingle
BenchmarkQueryEntitiesSingle-3                 8         140165732 ns/op
```

Adding bare entities is pretty slow off the bat, for something we may need to do hundreds of times per turn.
However, that's not even a realistic usecase: we will usually add an entity together with its components.

Two empty components makes it even slower, this will not scale well.

Checking for entity existence is slightly better but still not optimal.

Checking for component presence is not as bad as expected, unlikely we can do much better here.

Removing also much slower than we'd like it to be.

Querying however, which nearly every system does, is the real killer. Even a single component is not better,
as it needs to do almost the same amount of work in the worst case.

### After replacing the maps with slices + bitsets for querying:

```
BenchmarkMustGetEntityComponent
BenchmarkMustGetEntityComponent-3       449636858                2.760 ns/op           0 B/op          0 allocs/op
BenchmarkGetEntityComponent
BenchmarkGetEntityComponent-3           209203453                5.520 ns/op           0 B/op          0 allocs/op
BenchmarkSetEntityComponent
BenchmarkSetEntityComponent-3           50575398                23.66 ns/op            0 B/op          0 allocs/op
BenchmarkRemoveEntityComponent
BenchmarkRemoveEntityComponent-3        201453700                6.048 ns/op           0 B/op          0 allocs/op
BenchmarkAddEntity
BenchmarkAddEntity-3                    45559728                24.53 ns/op           45 B/op          0 allocs/op
BenchmarkHasEntity
BenchmarkHasEntity-3                    1000000000               1.014 ns/op           0 B/op          0 allocs/op
BenchmarkAddEntityComponent
BenchmarkAddEntityComponent-3            5057528               221.5 ns/op           233 B/op          0 allocs/op
BenchmarkHasEntityComponent
BenchmarkHasEntityComponent-3           295670764                4.165 ns/op           0 B/op          0 allocs/op
BenchmarkRemoveEntity
BenchmarkRemoveEntity-3                  3058321               347.0 ns/op             0 B/op          0 allocs/op
BenchmarkQueryEntitiesIter
BenchmarkQueryEntitiesIter-3                 757           1924979 ns/op         4136992 B/op          3 allocs/op
BenchmarkQueryEntitiesSingle
BenchmarkQueryEntitiesSingle-3            249214              4744 ns/op            6868 B/op          3 allocs/op
```

Definitely not bad, querying has become 2 orders of magnitude faster. Let's see if we can get a bit more out
by being smarter while adding entities, and a more optimized bitset package for our data sizes.