# rules for builing operators from skeleton
# operator-sdk new example-operator
#    --kind=ExampleService
#    --api-version=example.org/v2alpha1
#    --type=helm
#    --helm-chart ~/src/helm-operator-example/chart/

HelmOperator = provider(fields=["chart", "watches"])

def helm_operator(ctx):
  for f in ctx.attr.chart.files.to_list():
    if f.basename == "Chart.yaml":
      print(f.path)
      chart_path = f
      break

  out_dir = ctx.actions.declare_directory("operator")
  args = [
    "new",
    ctx.attr.name,
    "--kind=%s" % ctx.attr.kind,
    "--type=helm",
    "--helm-chart=${EXECROOT}/%s" % (chart_path.dirname),
  ]
  ctx.actions.run_shell(
    inputs = ctx.attr.chart.files,
    outputs = [out_dir],
    progress_message = "Generating operator-sdk against %s/%s" % (ctx.build_file_path, chart_path),
    command = "EXECROOT=$(pwd);cd " + out_dir.path +"; operator-sdk " + " ".join(args),
    use_default_shell_env = True,
  )
  return [DefaultInfo(files=depset([out_dir]))]

atom_helm_operator = rule(
  implementation = helm_operator,
  attrs = {
    "kind": attr.string(default="", doc="Kind to represented in the output schema (ex. ExampleApp)", mandatory=True),
    "version": attr.string(default="", doc="Version of the API (ex. example.org/v1)", mandatory=True),
    "chart": attr.label(default=":chart", doc="Location helm chart to run operator-sdk against", mandatory=True),
  }
)