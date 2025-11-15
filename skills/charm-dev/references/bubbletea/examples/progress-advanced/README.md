# Progress Advanced

An advanced progress bar example showcasing sophisticated gradient techniques. This demonstrates what's possible with the current API and points to upcoming features from [bubbles PR #838](https://github.com/charmbracelet/bubbles/pull/838).

## What This Example Shows

This example demonstrates four different gradient configurations:

### 1. Warm Gradient (Red → Yellow)
- Perfect for error, warning, or alert states
- Smooth transition from red (#FF0000) to yellow (#FFFF00)
- The gradient spans the full width of the bar

### 2. Cool Gradient (Blue → Cyan)
- Great for info messages or processing states
- Smooth transition from blue (#0000FF) to cyan (#00FFFF)
- Professional, calm appearance

### 3. Scaled Gradient (Green shades)
- Uses `WithScaledGradient()` instead of `WithGradient()`
- The gradient **always scales to fit the filled portion**
- You always see the full gradient, regardless of progress
- Particularly useful when you want consistent color distribution

### 4. Vibrant Gradient (Default)
- Uses the default purple-to-pink gradient
- Beautiful out-of-the-box appearance
- Shows what you get with `WithDefaultGradient()`

## About PR #838: ColorFunc

The upcoming `ColorFunc` feature ([PR #838](https://github.com/charmbracelet/bubbles/pull/838)) will enable even more advanced progress bar styling:

### What ColorFunc Enables:
- **Dynamic Color Changes**: Change colors based on progress milestones (e.g., red → yellow → green)
- **Animated Patterns**: Create striped effects that move as the bar fills
- **Custom Algorithms**: Use HSL, time-based animations, or any custom color logic
- **Position-Aware Coloring**: Color each character based on both total progress and position

### ColorFunc API (Coming Soon):
```go
type ColorFunc func(total, current float64) color.Color
```

Where:
- `total`: Overall completion percentage (0.0 to 1.0)
- `current`: Position within the bar being colored (0.0 to 1.0)

### Example Use Cases:
```go
// Milestone-based colors
progress.WithColorFunc(func(total, current float64) color.Color {
    if total <= 0.33 { return red }
    if total <= 0.66 { return yellow }
    return green
})

// Animated stripes
progress.WithColorFunc(func(total, current float64) color.Color {
    if int(current*10)%2 == 0 { return blue1 }
    return blue2
})

// HSL rainbow
progress.WithColorFunc(func(total, current float64) color.Color {
    hue := current * 360.0
    return hslToRGB(hue, 100, 50)
})
```

## Running the Example

```bash
go run main.go
```

The four progress bars will animate from 0% to 100%, demonstrating different gradient styles.

## Related

- [Bubbles Progress Component](https://github.com/charmbracelet/bubbles/tree/master/progress)
- [PR #838: ColorFunc Feature](https://github.com/charmbracelet/bubbles/pull/838)
- [Charm Bubbles Documentation](https://github.com/charmbracelet/bubbles)
