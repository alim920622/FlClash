import 'package:fl_clash/common/common.dart';
import 'package:fl_clash/models/models.dart';
import 'package:fl_clash/widgets/widgets.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class NetworkSpeed extends StatefulWidget {
  const NetworkSpeed({super.key});

  @override
  State<NetworkSpeed> createState() => _NetworkSpeedState();
}

class _NetworkSpeedState extends State<NetworkSpeed> {
  List<Point> initPoints = const [Point(0, 0), Point(1, 0)];

  List<Point> _getPoints(List<Traffic> traffics) {
    List<Point> trafficPoints = traffics
        .toList()
        .asMap()
        .map(
          (index, e) => MapEntry(
            index,
            Point(
              (index + initPoints.length).toDouble(),
              e.speed.toDouble(),
            ),
          ),
        )
        .values
        .toList();

    return [...initPoints, ...trafficPoints];
  }

  Traffic _getLastTraffic(List<Traffic> traffics) {
    if (traffics.isEmpty) return Traffic();
    return traffics.last;
  }

  @override
  Widget build(BuildContext context) {
    final color = context.colorScheme.onSurfaceVariant.toLight;
    return SizedBox(
      height: getWidgetHeight(2),
      child: CommonCard(
        onPressed: () {},
        info: Info(
          label: appLocalizations.networkSpeed,
          iconData: Icons.speed_sharp,
        ),
        child: Selector<AppFlowingState, List<Traffic>>(
          selector: (_, appFlowingState) => appFlowingState.traffics,
          builder: (_, traffics, __) {
            return Stack(
              children: [
                Positioned.fill(
                  child: Padding(
                    padding: EdgeInsets.all(16).copyWith(
                      bottom: 0,
                      left: 0,
                      right: 0,
                    ),
                    child: LineChart(
                      gradient: true,
                      color: Theme.of(context).colorScheme.primary,
                      points: _getPoints(traffics),
                    ),
                  ),
                ),
                Positioned(
                  top: 0,
                  right: 0,
                  child: Transform.translate(
                    offset: Offset(
                      -16,
                      -26,
                    ),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.end,
                      children: [
                        Text(
                          "${appLocalizations.upload} ${_getLastTraffic(traffics).up}/s",
                          style: context.textTheme.bodySmall?.copyWith(
                            color: color,
                          ),
                        ),
                        SizedBox(
                          height: 4,
                        ),
                        Text(
                          "${appLocalizations.download} ${_getLastTraffic(traffics).down}/s",
                          style: context.textTheme.bodySmall?.copyWith(
                            color: color,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
