import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;

//START OMIT
public class MainActivity extends AppCompatActivity {

    go.gobindex.Gobindex.Gopher mGopher = null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        // load Go and initialize the Go runtime.
        go.Go.init(this);

        // call like a Java function.
        mGopher = go.gobindex.Gobindex.NewGopher("Go Gopher #1");

        // use like a Java object.
        go.gobindex.Gobindex.Hello(mGopher);

        // when the last reference finalizes, the corresponding
        // Go object will be unpinned.
    }
//END OMIT

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }
...
}

